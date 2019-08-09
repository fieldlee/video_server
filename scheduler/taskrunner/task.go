package taskrunner

import (
	"errors"
	"log"
	"os"
	"sync"
	"video_server/scheduler/dbops"
)

func deleteVideoFile(vid string) error {
	err := os.Remove("./videos/"+vid)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

func VideoClearDispatcher(dc dataChan) error{
	res,err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("VideoClearDispatcher:%v",err)
		return err
	}
	if len(res) == 0{
		return errors.New("task finished")
	}
	for _,id := range res{
		dc <- id
	}
	return nil
}

func VideoClearExecuter(dc dataChan) error{
	errMsp := &sync.Map{}
	var lastErr error
	forloop:
		for{
			select {
			case id := <- dc:
				go func(vid string) {
					err := deleteVideoFile(vid)
					if err != nil {
						errMsp.Store(vid,err)
						return
					}
					err = dbops.DeleteVideoDeletionRecord(vid)
					if err != nil {
						errMsp.Store(vid,err)
						return
					}
				}(id.(string))

			default:
				break forloop
			}
		}
	//获得msp里的err信息
	errMsp.Range(func(k, v interface{}) bool {
		lastErr = v.(error)
		if lastErr != nil {
			return false
		}
		return true
	})
	return lastErr
}