package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {

	var (
		putResp               *clientv3.PutResponse
		getResp               *clientv3.GetResponse
		delResp               *clientv3.DeleteResponse
		leaseResp, leaseResp1 *clientv3.LeaseGrantResponse
	)

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connection field !", err)
		return
	}

	fmt.Println("connection seccuss")
	fmt.Println(client.Password, client.Password)
	defer client.Close()

	//ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	//
	//_, err = client.Put(ctx, "logagent/conf", "sample_value")
	//cancelFunc()
	//if err != nil {
	//	fmt.Println("put failed , err ", err)
	//	return
	//}
	//ctx, cancelFunc = context.WithTimeout(context.Background(), time.Second)
	////response, err := client.Get(ctx, "logagent/conf")
	//cancelFunc()
	//
	//if err != nil {
	//	fmt.Println("get field ,err", err)
	//	return
	//}

	//for _, v := range response.Kvs {
	//	fmt.Printf("varsion is %d, %s : %s\n", v.Version, v.Key, v.Value)
	//}
	//

	if putResp, err = client.Put(context.TODO(), "logagent/conf", "88888888"); err != nil {
		return
	} else {
		fmt.Println(putResp.Header)
		fmt.Println(putResp.PrevKv)
	}
	//
	//for {
	//	watchs := client.Watch(context.Background(), "logagent/conf")
	//
	//	for v := range watchs {
	//		for _, v := range v.Events {
	//			fmt.Printf("%s %q :%q \n", v.Type, v.Kv.Key, v.Kv.Value)
	//		}
	//	}
	//}

	kv := clientv3.NewKV(client)
	if putResp, err = kv.Put(context.TODO(), "/cron/jobs/job2", "hello", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
		return
	} else {
		if putResp.PrevKv != nil {
			fmt.Println(string(putResp.PrevKv.Key), string(putResp.PrevKv.Value))
		}
	}

	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/", clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
		return
	} else {

		fmt.Println(getResp.Kvs)
	}

	if delResp, err = kv.Delete(context.TODO(), "/cron/jobs/job2", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(len(delResp.PrevKvs))
	}

	if leaseResp, err = client.Grant(context.TODO(), 10); err != nil {
		fmt.Println(err)
		return
	}

	if putResp, err = client.Put(context.TODO(), "/cron/lock/job1", "ok", clientv3.WithLease(leaseResp.ID)); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(putResp.Header.Revision)
	}

	if keepAliveChan, err := client.KeepAlive(context.TODO(), leaseResp.ID); err != nil {
		fmt.Println(err)
		return
	} else {
		go func() {
			for {
				select {
				case resp := <-keepAliveChan:
					if resp == nil {
						fmt.Println("续租失败")
						goto END
					} else {
						fmt.Println("续租成功")
					}
				}
			}
		END:
		}()
	}
	k := 8
	for k != 0 {
		if getResp, err = client.Get(context.TODO(), "/cron/lock/job1"); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println(getResp.Count)
			time.Sleep(3 * time.Second)
		}
		k--
	}

	go func() {
		for {
			client.Put(context.TODO(), "/cron/watch/job1", "I am watch job1")
			time.Sleep(1000)
			client.Delete(context.TODO(), "/cron/watch/job1")
			time.Sleep(1000)
		}
	}()
}
