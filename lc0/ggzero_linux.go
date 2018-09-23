// A new client to work with the lc0 binary.
//
//

package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/nightlyone/lockfile"
	"io"
	"io/ioutil"

	//"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"client"

	// "github.com/Tilps/chess" 这儿得用中国的
)




// Settings holds username and password.
type Settings struct {
	User string
	Pass string
	Args string

	//IsDownloading bool    			// 当前是不是在下载权重了，如果是，这个界面就不要下载了，只要等待完成
	IsSaveTrain bool
	//isSaveMatch bool
}

var (
	startTime  time.Time
	totalGames int

	//hostname = flag.String("hostname", "http://lcc.ya.cn", "Address of the server")
	hostname = flag.String("hostname", "http://47.104.173.38", "Address of the server")
	//hostname = flag.String("hostname", "http://47.104.173.38", "Address of the server")

    dls_hostname = flag.String("hostname2", "http://144.202.54.34", "server from usa dls for colab")

	user     = flag.String("user", "", "Username")
	password = flag.String("password", "", "Password")
	//gpu      = flag.Int("gpu", -1, "ID of the OpenCL device to use (-1 for default, or no GPU)")
	debug    = flag.Bool("debug", false, "Enable debug mode to see verbose output and save logs")
	lc0Args  = flag.String("lc0args", "", `Extra args to pass to the backend.  example: --lc0args="--parallelism=10 --threads=2"`)
)

var SettingFile = "train.json"
var VERSION  = "v0.18"
var settings = Settings{}
var Ctrain cmdWrapper   // 全局变量，这个是

/*
	Reads the user and password from a config file and returns empty strings if anything went wrong.
	If the config file does not exists, it prompts the user for a username and password and creates the config file.
*/

func SaveSettings(path string, myuser string, mypassword string){

	file, err := os.Open(path)
	if err != nil {
		// File was not found
		//fmt.Printf("GGzeor Cuda训练客户端 Version: %s.\n", VERSION)
		//fmt.Printf("请输入用户名与口令，会自动建立一个帐号.\n")
		//fmt.Printf("请注意，口令将以明码保存，因此请不要输入您用于敏感场合的保密的口令！\n")
		//fmt.Printf("请妥善保管您的口令，此口令暂时不可恢复！\n")
		//fmt.Printf("请输入用户名 : ")
		//fmt.Scanf("%s\n", &settings.User)
		//fmt.Printf("请输入口令 : ")
		//fmt.Scanf("%s\n", &settings.Pass)

		settings.User = myuser;
		settings.Pass = mypassword;

		// 缺省的数据
		settings.Args = "--threads=1 --minibatch-size=1024 --backend=multiplexing --backend-opts=(backend=cudnn,gpu=0) --parallelism=8 --nncache=800000"
		//
		//settings.DeleteOldWeightsWhenOpen = false
		//settings.DeleteTrainWhenOpen = false
		//settings.IsDownloading = false
		settings.IsSaveTrain = false          //
		//settings.isSaveMatch = false          //
		//settings.IsDownFinished = false

		jsonSettings, err := json.Marshal(settings)
		if err != nil {
			log.Fatal("Cannot encode settings to JSON ", err)
			return;// "", ""
		}
		settingsFile, err := os.Create(path)
		defer settingsFile.Close()
		if err != nil {
			log.Fatal("Could not create output file ", err)
			return; // "", ""
		}
		settingsFile.Write(jsonSettings)
		return; // settings.User, settings.Pass
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&settings)
	if err != nil {
		log.Fatal("Error decoding JSON ", err)
		return; // "", ""
	}

	// 打开界面，清空旧数据
	//if settings.DeleteTrainWhenOpen == true{

	//}

	// 看一下是不是要删除旧的权重,删除一天前的权重
	//if settings.DeleteOldWeightsWhenOpen == true{
	//	now_time := time.Now().Unix()
	//	err := filepath.Walk("networks")
	//
	//}

	return; // settings.User, settings.Pass
}

func readSettings(path string) (string, string) {
	// settings := Settings{}

	file, err := os.Open(path)
	if err != nil {
		// File was not found
		fmt.Printf("GGzeor Cuda训练客户端 Version: %s.\n", VERSION)
		fmt.Printf("请输入用户名与口令，会自动建立一个帐号.\n")
		fmt.Printf("请注意，口令将以明码保存，因此请不要输入您用于敏感场合的保密的口令！\n")
		fmt.Printf("请妥善保管您的口令，此口令暂时不可恢复！\n")
		fmt.Printf("请输入用户名 : ")
		fmt.Scanf("%s\n", &settings.User)
		fmt.Printf("请输入口令 : ")
		fmt.Scanf("%s\n", &settings.Pass)

		// 缺省的数据
		settings.Args = "--threads=1 --minibatch-size=1024 --backend=multiplexing --backend-opts=(backend=cudnn,gpu=0) --parallelism=8 --nncache=800000"
		//
		//settings.DeleteOldWeightsWhenOpen = false
		//settings.DeleteTrainWhenOpen = false
		//settings.IsDownloading = false
		settings.IsSaveTrain = false          //
		//settings.isSaveMatch = false          //
		//settings.IsDownFinished = false

		jsonSettings, err := json.Marshal(settings)
		if err != nil {
			log.Fatal("Cannot encode settings to JSON ", err)
			return "", ""
		}
		settingsFile, err := os.Create(path)
		defer settingsFile.Close()
		if err != nil {
			log.Fatal("Could not create output file ", err)
			return "", ""
		}
		settingsFile.Write(jsonSettings)
		return settings.User, settings.Pass
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&settings)
	if err != nil {
		log.Fatal("Error decoding JSON ", err)
		return "", ""
	}

	// 打开界面，清空旧数据
	//if settings.DeleteTrainWhenOpen == true{

    //}

	// 看一下是不是要删除旧的权重,删除一天前的权重
	//if settings.DeleteOldWeightsWhenOpen == true{
	//	now_time := time.Now().Unix()
	//	err := filepath.Walk("networks")
	//
	//}

	return settings.User, settings.Pass
}

func getExtraParams() map[string]string {
	return map[string]string{
		"user":     *user,
		"password": *password,
		"version":  "18",
	}
}

func uploadGame(httpClient *http.Client, path string, pgn string,
	nextGame client.NextGameResponse, version string) error {

	var retryCount uint32

	for {
		retryCount++
		if retryCount > 3 {
			return errors.New("上传训练数据失败: Too many retries")
		}

		extraParams := getExtraParams()
		extraParams["training_id"] = strconv.Itoa(int(nextGame.TrainingId))
		extraParams["network_id"] = strconv.Itoa(int(nextGame.NetworkId))
		extraParams["pgn"] = pgn
		extraParams["engineVersion"] = version

		request, err := client.BuildUploadRequest(*hostname+"/upload_game", extraParams, "file", path)
		if err != nil {
			log.Printf("BUR: %v", err)
			return err
		}
		resp, err := httpClient.Do(request)
		if err != nil {
			log.Printf("http.Do: %v", err)
			return err
		}
		body := &bytes.Buffer{}
		_, err = body.ReadFrom(resp.Body)
		if err != nil {
			log.Print(err)
			log.Print("上传出错了, 正在重试...")
			time.Sleep(time.Second * (2 << retryCount))
			continue
		}
		resp.Body.Close()
		break
	}

	totalGames++
	log.Printf("成功上传了 %d 局棋局 于 %s 时间内", totalGames, time.Since(startTime))

	err := os.Remove(path)
	if err != nil {
		log.Printf("删除训练文件失败了: %v", err)
	}

	return nil
}

type gameInfo struct {
	pgn   string
	fname string
}

type cmdWrapper struct {
	Cmd      *exec.Cmd
	Pgn      string
	Input    io.WriteCloser
	BestMove chan string
	gi       chan gameInfo
	Version  string
	result string
	isTrainging bool
}

func (c *cmdWrapper) openInput() {
	var err error
	c.Input, err = c.Cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
}

func convertMovesToPGN(moves []string) string {
	//game := chess.NewGame(chess.UseNotation(chess.LongAlgebraicNotation{}))
	//for _, m := range moves {
	//	err := game.MoveStr(m)
	//	if err != nil {
	//		log.Fatalf("movstr: %v", err)
	//	}
	//}
	//game2 := chess.NewGame()
	//b, err := game.MarshalText()
	//if err != nil {
	//	log.Fatalf("MarshalText failed: %v", err)
	//}
	//game2.UnmarshalText(b)
	//return game2.String()
	var pgn string
	for _, m := range moves {
		pgn = pgn + m + " "
	}
	return pgn
}

func createCmdWrapper() *cmdWrapper {
	c := &cmdWrapper{
		gi:       make(chan gameInfo),
		BestMove: make(chan string),
	}
	return c
}

func (c *cmdWrapper) launchMatch(networkPath string, args []string, input bool) {
	dir, _ := os.Getwd()
	c.Cmd = exec.Command(path.Join(dir, "lc0")) // 程序文件名
	c.Cmd.Args = append(c.Cmd.Args, args...)
	// c.Cmd.Args = append(c.Cmd.Args, fmt.Sprintf("--weights=%s", networkPath))
	if *lc0Args != "" {
		// TODO: We might want to inspect these to prevent someone
		// from passing a different visits or batch size for example.
		// Possibly just exposts exact args we want to passthrough like
		// backend.  For testing right now this is probably useful to not
		// need to rebuild the client to change lc0 args.
		parts := strings.Split(*lc0Args, " ")
		c.Cmd.Args = append(c.Cmd.Args, parts...)
	}
	if !*debug {
		//		c.Cmd.Args = append(c.Cmd.Args, "--quiet")
		fmt.Println("lc0 is never quiet.")
	}
	fmt.Printf("Args: %v\n", c.Cmd.Args)

	stdout, err := c.Cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	stderr, err := c.Cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer close(c.BestMove)
		defer close(c.gi)
		stdoutScanner := bufio.NewScanner(stdout)
		for stdoutScanner.Scan() {
			line := stdoutScanner.Text()
			switch {
			case strings.HasPrefix(line, "gameready"):
				parts := strings.Split(line, " ")
				if 	parts[1] != "gameid" ||
					parts[3] != "result" ||
					parts[5] != "pgn" {
					log.Printf("Malformed gameready: %q", line)
					break
				}
				// file := ""
				//gameid := parts[4]
				//result := parts[8]
				// pgn := convertMovesToPGN(parts[10:])
				c.result = parts[4]   // 得到比赛结果
				pgn := convertMovesToPGN(parts[6:])
				c.Pgn = pgn
				//fmt.Printf("PGN: %s\n", pgn)
				//c.gi <- gameInfo{pgn: pgn, fname: file}
			//case strings.HasPrefix(line, "bestmove "):
			//	fmt.Println(line)
			//	c.BestMove <- strings.Split(line, " ")[1]
			case strings.HasPrefix(line, "id name lczero "):
				c.Version = strings.Split(line, " ")[3]
			//case strings.HasPrefix(line, "info"):
			//	break
			//	fallthrough
			default:
				fmt.Println(line)
			}
		}
	}()

	go func() {
		stderrScanner := bufio.NewScanner(stderr)
		for stderrScanner.Scan() {
			fmt.Printf("%s\n", stderrScanner.Text())
		}
	}()

	if input {
		c.openInput()
	}

	err = c.Cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func (c *cmdWrapper) launch(networkPath string, args []string, input bool) {
	dir, _ := os.Getwd()
	c.Cmd = exec.Command(path.Join(dir, "lc0"))  // 程序文件名
	c.Cmd.Args = append(c.Cmd.Args, args...)
	c.Cmd.Args = append(c.Cmd.Args, fmt.Sprintf("--weights=%s", networkPath))
	if *lc0Args != "" {
		// TODO: We might want to inspect these to prevent someone
		// from passing a different visits or batch size for example.
		// Possibly just exposts exact args we want to passthrough like
		// backend.  For testing right now this is probably useful to not
		// need to rebuild the client to change lc0 args.
		parts := strings.Split(*lc0Args, " ")
		c.Cmd.Args = append(c.Cmd.Args, parts...)
	}
	if !*debug {
		//		c.Cmd.Args = append(c.Cmd.Args, "--quiet")
		fmt.Println("lc0 is never quiet.")
	}
	fmt.Printf("Args: %v\n", c.Cmd.Args)

	stdout, err := c.Cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	stderr, err := c.Cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer close(c.BestMove)
		defer close(c.gi)
		stdoutScanner := bufio.NewScanner(stdout)
		for stdoutScanner.Scan() {
			line := stdoutScanner.Text()
			switch {
			case strings.HasPrefix(line, "gameready"):
				parts := strings.Split(line, " ")
				if parts[1] != "trainingfile" ||
					parts[3] != "gameid" ||
					parts[5] != "player1" ||
					parts[7] != "result" ||
					parts[9] != "pgn" {
					log.Printf("Malformed gameready: %q", line)
					break
				}
				file := parts[2]
				//gameid := parts[4]
				//result := parts[8]
				pgn := convertMovesToPGN(parts[10:])
				//pgn := parts[9];
				//fmt.Printf("PGN: %s\n", pgn)
				c.gi <- gameInfo{pgn: pgn, fname: file}
			case strings.HasPrefix(line, "bestmove "):
				fmt.Println(line)
				c.BestMove <- strings.Split(line, " ")[1]
			case strings.HasPrefix(line, "id name lczero "):
				c.Version = strings.Split(line, " ")[3]
			case strings.HasPrefix(line, "info"):
				break
				fallthrough
			default:
				fmt.Println(line)
			}
		}
	}()

	go func() {
		stderrScanner := bufio.NewScanner(stderr)
		for stderrScanner.Scan() {
			fmt.Printf("%s\n", stderrScanner.Text())
		}
	}()

	if input {
		c.openInput()
	}

	// 当前正在训练了
	Ctrain = *c
	Ctrain.isTrainging = true    //

	err = c.Cmd.Start()
	if err != nil {
		//log.Fatal(err)
		err = nil
	}
}

func playMatch(baselinePath string, candidatePath string, params []string, flip bool) (int, string, string, error) {

	train_cmd := fmt.Sprintf("matchplay")
	params = append(params, train_cmd)

	train_cmd = fmt.Sprintf("--baselinePath=%s", baselinePath)
	params = append(params, train_cmd)

	train_cmd = fmt.Sprintf("--candidatePath=%s", candidatePath)
	params = append(params, train_cmd)

	train_cmd = fmt.Sprintf("--parallelism=1")  // 一个线程就行了
	params = append(params, train_cmd)

	train_cmd = fmt.Sprintf("--nodes=800")      // 比赛参数
	params = append(params, train_cmd)

	train_cmd = fmt.Sprintf("--games=1")      // // 只比赛一局
	params = append(params, train_cmd)

	ArgArray := strings.Fields(strings.TrimSpace(settings.Args))
	for _, str := range ArgArray {
		params = append(params, str)
	}
	//params = append(params,fmt.Sprintf("--gpu=%v", settings.GPU_No))
	//params = append(params,fmt.Sprintf("--threads=%v", settings.Thread))

	var flipCmd string
	if flip {
		flipCmd = fmt.Sprintf("--flip=%d", 1)
	} else{
		flipCmd = fmt.Sprintf("--flip=%d", 0)
	}
	params = append(params, flipCmd)

	c := createCmdWrapper()
	c.Version = VERSION
	c.launchMatch(baselinePath, params, false)

	err := c.Cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}

	ires := 0
	if c.result == "whitewon"{
		ires = -1
	}	else if c.result == "draw"{
		ires = 0
	}	else if c.result == "blackwon" {
		ires = 1
	} else{
		c.result = "draw"
		fmt.Printf("**********************  match result error:\n")
		fmt.Printf("match result error: result %s\n", c.result)
		fmt.Printf("fen: %s\n", c.Pgn)
		fmt.Printf("********************** 请复制上面的信息:\n")

		c.result = "draw"
		//log.Fatal("match result error ! ")
	}

	// Always report the result relative to the candidate engine
	// (which defaults to white, unless flip = true)
	if flip {
		ires = -ires
	}

	// chess.UseNotation(chess.AlgebraicNotation{})(game)
	//fmt.Printf("PGN: %s\n", c.Pgn) // fmt.Printf("PGN: %s\n", game.String())
	return ires, c.Pgn, c.Version, nil
}

func train(httpClient *http.Client, ngr client.NextGameResponse,
	networkPath string, count int, params []string, doneCh chan bool) {
	// pid is intended for use in multi-threaded training
	pid := os.Getpid()

	dir, _ := os.Getwd()
	if *debug {
		logsDir := path.Join(dir, fmt.Sprintf("logs-%v", pid))
		os.MkdirAll(logsDir, os.ModePerm)
		logfile := path.Join(logsDir, fmt.Sprintf("%s.log", time.Now().Format("20060102150405")))
		params = append(params, "-l"+logfile)
	}

	// lc0 needs selfplay first in the argument list.
	params = append([]string{"selfplay"}, params...)
	params = append(params, "--training=true")

	// resign
	//params = append(params,"--resign-percentage=4.0")
	//params = append(params,"--resign-playthrough=10")

	ArgArray := strings.Fields(strings.TrimSpace(settings.Args))
	for _, str := range ArgArray {
		params = append(params, str)
	}

	if settings.IsSaveTrain{
		params = append(params,"--IsSaveTrainPgn=true")
	} else{
		params = append(params,"--IsSaveTrainPgn=false")
	}

	// params = append(params, "--tempdecay-moves=10")  // BY nooby

	//arams = append(params,fmt.Sprintf("--gpu=%v", settings.GPU_No))
	//params = append(params,fmt.Sprintf("--threads=%v", settings.Thread))

	c := createCmdWrapper()
	c.Version = VERSION
	c.launch(networkPath, params /* input= */, false)
	for done := false; !done; {
		numGames := 1
		select {
		case <-doneCh:
			done = true
			log.Println("Received message to end training, killing lc0")
			c.Cmd.Process.Kill()
		case _, ok := <-c.BestMove:
			// Just swallow the best moves, only needed for match play.
			if !ok {
				log.Printf("BestMove channel closed unexpectedly, exiting train loop")
				break
			}
		case gi, ok := <-c.gi:
			if !ok {
				log.Printf("GameInfo channel closed, exiting train loop")
				done = true
				break
			}
			fmt.Printf("Uploading game: %d\n", numGames)
			numGames++
			go uploadGame(httpClient, gi.fname, gi.pgn, ngr, c.Version)
		}
	}

	log.Println("Waiting for lc0 to stop")
	err := c.Cmd.Wait()
	if err != nil {
		//log.Fatal(err)
		err = nil
	}
	log.Println("lc0 stopped")
}


func checkValidNetwork(dir string, sha string) (string, error) {
	// Sha already exists?
	path := filepath.Join(dir, sha)
	_, err := os.Stat(path)
	if err == nil {
		file, _ := os.Open(path)
		reader, err := gzip.NewReader(file)
		if err == nil {
			_, err = ioutil.ReadAll(reader)
		}
		file.Close()
		if err != nil {
			fmt.Printf("Deleting invalid network...\n")
			os.Remove(path)
			return path, err
		} else {
			return path, nil
		}
	}
	return path, err
}


func removeAllExcept(dir string, sha string) (error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.Name() == sha {
			continue
		}
		fmt.Printf("Removing %v\n", file.Name())
		err := os.RemoveAll(filepath.Join(dir, file.Name()))
		if err != nil {
			return err
		}
	}
	return nil
}

func acquireLock(dir string, sha string) (lockfile.Lockfile, error) {
	lockpath, _ := filepath.Abs(filepath.Join(dir, sha + ".lck"))
	lock, err := lockfile.New(lockpath)
	if err != nil {
		// Unknown error. Exit.
		log.Fatalf("Cannot init lockfile: %v", err)
	}
	// Attempt to acquire lock
	err = lock.TryLock()
	return lock, err
}

func getNetwork(httpClient *http.Client, sha string, clearOld bool) (string, error) {
	dir := "networks"
	os.MkdirAll(dir, os.ModePerm)
	path, err := checkValidNetwork(dir, sha)
	if err == nil {
		// There is already a valid network. Use it.
		if clearOld {
			err := removeAllExcept(dir, sha)
			if err != nil {
				log.Printf("Failed to remove old network(s): %v", err)
			}
		}
		return path, nil
	}

	// Otherwise, let's download it
	lock, err := acquireLock(dir, sha)

	if err != nil {
		if err == lockfile.ErrBusy {
			log.Println("Download initiated by other client")
			return "", err
		} else {
			log.Fatalf("Unable to lock: %v", err)
		}
	}

	// Lockfile acquired, download it
	defer lock.Unlock()
	fmt.Println("Downloading network...")
	err = client.DownloadNetwork(httpClient, *hostname, path, sha)
	if err != nil {
		log.Printf("Network download failed: %v", err)
		return "", err
	}
	return checkValidNetwork(dir, sha)
}

func validateParams(args []string) []string {
	validArgs := []string{}
	for _, arg := range args {
		if strings.HasPrefix(arg, "--tempdecay") {
			continue
		}
		validArgs = append(validArgs, arg)
	}
	return validArgs
}

// goon train
//var Train_nextGame client.NextGameResponse
//var Train_serverParams []string

func nextGame(httpClient *http.Client, count int) error {
	nextGame, err := client.NextGame(httpClient, *hostname, getExtraParams())
	if err != nil {
		return err
	}
	var serverParams []string
	err = json.Unmarshal([]byte(nextGame.Params), &serverParams)
	if err != nil {
		return err
	}
	log.Printf("服务器参数: %s", serverParams)
	serverParams = validateParams(serverParams)

	if nextGame.Type == "match" {

		log.Printf("比赛开始了。。。 flip= %t", nextGame.Flip)

		log.Println("\nLinux 版本暂时不能比赛 本窗口等待40秒")
		time.Sleep(time.Second*40)
		return nil

		/*
		networkPath, err := getNetwork(httpClient, nextGame.Sha, false)
		if err != nil {
			return err
		}
		candidatePath, err := getNetwork(httpClient, nextGame.CandidateSha, false)
		if err != nil {
			return err
		}
		*/


		/*
		result, pgn, version, err := playMatch(networkPath, candidatePath, serverParams, nextGame.Flip)

		if err != nil {
			log.Fatalf("比赛出错: %v", err)
			return err
		}
		extraParams := getExtraParams()
		extraParams["engineVersion"] = version
		log.Println("uploading match result")
		go client.UploadMatchResult(httpClient, *hostname, nextGame.MatchGameId, result, pgn, extraParams)

		*/

		/*

		//fmt.Printf("\nLinux 版本暂时不能比赛 本窗口等待20秒")
		// time.Sleep(time.Second*20)

		if len(Train_serverParams) < 1{
			fmt.Printf("\nLinux 版本比赛时没有得到训练数据，等待20秒！")
			time.Sleep(time.Second*20)
			return nil
		}

		fmt.Printf("\nLinux 版本比赛时得到训练数据，继续训练中。。！")

		nextGame = Train_nextGame
		serverParams = Train_serverParams

		networkPath, err := getNetwork(httpClient, nextGame.Sha, false)
		if err != nil {
			return err
		}
		doneCh := make(chan bool)
		go func() {
			errCount := 0
			for {
				time.Sleep(60 * time.Second)
				ng, err := client.NextGame(httpClient, *hostname, getExtraParams())
				if err != nil {
					fmt.Printf("Error talking to server: %v\n", err)
					errCount++
					if errCount < 10 {
						continue
					}
				}
				// by lgl
				// ng.Type = "match"
				// by lgl
				if err != nil || ng.Type != nextGame.Type || ng.Sha != nextGame.Sha {
					doneCh <- true
					close(doneCh)
					return
				}
				errCount = 0
			}
		}()
		train(httpClient, nextGame, networkPath, count, serverParams, doneCh)
		//train(httpClient, nextGame, networkPath, count, []string{"--visits=800"}, doneCh)
		return nil
		*/
	}



	if nextGame.Type == "train" {

		// 保存训练的参数
		//Train_nextGame = nextGame
		//Train_serverParams = serverParams

		networkPath, err := getNetwork(httpClient, nextGame.Sha, false)
		if err != nil {
			return err
		}
		doneCh := make(chan bool)
		go func() {
			errCount := 0
			for {
				time.Sleep(60 * time.Second)
				ng, err := client.NextGame(httpClient, *hostname, getExtraParams())
				if err != nil {
					fmt.Printf("Error talking to server: %v\n", err)
					errCount++
					if errCount < 10 {
						continue
					}
				}
				// by lgl
				// ng.Type = "match"
				// by lgl
				if err != nil || ng.Type != nextGame.Type || ng.Sha != nextGame.Sha {
					doneCh <- true
					close(doneCh)
					return
				}
				errCount = 0
			}
		}()
		train(httpClient, nextGame, networkPath, count, serverParams, doneCh)
		//train(httpClient, nextGame, networkPath, count, []string{"--visits=800"}, doneCh)
		return nil
	}

	return errors.New("Unknown game type: " + nextGame.Type)
}

func main() {



		flag.Parse()

		if len(*user) == 0 || len(*password) == 0 {
			*user, *password = readSettings(SettingFile)
		}	else{
			SaveSettings(SettingFile, *user, *password)
		}

		if len(*user) == 0 {
			log.Fatal("You must specify a username")
		}
		if len(*password) == 0 {
			log.Fatal("You must specify a non-empty password")
		}

		httpClient := &http.Client{}
		startTime = time.Now()
		for i := 0; ; i++ {
			err := nextGame(httpClient, i)
			if err != nil {
				log.Print(err)
				log.Print("Sleeping for 30 seconds...")
				time.Sleep(30 * time.Second)
				continue
			}
		}

}
