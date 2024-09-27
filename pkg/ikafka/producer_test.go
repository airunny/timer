package ikafka

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/IBM/sarama"
	"github.com/stretchr/testify/assert"
)

type Message struct {
	Index int `json:"index"`
}

func (e *Message) Encoder() sarama.StringEncoder {
	str, _ := json.Marshal(e)
	return sarama.StringEncoder(str)
}

type AppFlyer struct {
	EventTime   string `json:"event_time"`
	EventName   string `json:"event_name"`
	Platform    string `json:"platform"`
	AppsflyerId string `json:"appsflyer_id"`
	AppName     string `json:"app_name"`
}

func TestNewProducer3(t *testing.T) {
	producer, err := NewProducer(&ProducerConfig{
		Brokers: []string{"127.0.0.1:9092", "127.0.0.1:9093", "127.0.0.1:9094"},
		Version: "3.6.0",
	})
	assert.Nil(t, err)

	wg := &sync.WaitGroup{}
	wg.Add(4)
	go func() {
		for i := 0; i < 500000; i++ {
			// TODO
			msg := sarama.StringEncoder(`{"api_version":"2.0","media_source":"organic","wifi":false,"country_code":"CN","event_type":"organic-install-in-app-event","sdk_version":"v6.12.1","city":"Beijing","app_id":"com.foreigncurrency.internationalfxeye","event_time":"2024-01-29 04:01:42","tags":["k8s_business","ACK-SH","beats_input_codec_json_applied"],"app_name":"WikiFX-Broker Regulatory APP","@timestamp":"2024-01-28T20:01:45.138Z","dma":"156001","install_app_store":"ziji_qudao","device_download_time_selected_timezone":"2023-12-12 19:26:08.964+0800","install_time_selected_timezone":"2023-12-12 19:26:08.964+0800","platform":"android","BasicData_App":"3","user_agent":"Dalvik/2.1.0 (Linux; U; Android 10; Redmi Note 7 Pro Build/PKQ1.181203.001)","selected_currency":"CNY","region":"AS","device_download_time":"2023-12-12 11:26:08.964","event_detail":{"url":"https://wikidc.interface004.com/fxapp/fxcfg.json","uid":"2988393600","publicProperty":"{\"userId\":\"2988393600\",\"deviceId\":\"1,29,3,302,0,4de318f461e57002,0\",\"languageCode\":\"zh-CN\",\"countryCode\":\"156\",\"env\":\"正式环境\",\"onlyid\":\"4de318f461e57002\",\"isv3\":\"旧V3\"}"},"device_model":"xiaomi::Redmi Note 7 Pro","event_revenue_currency":"USD","is_retargeting":false,"device_category":"phone","bundle_id":"com.foreigncurrency.internationalfxeye","install_time":"2023-12-12 19:26:08","ip":"111.63.79.184","agent":{"id":"58f1fd7c-ea24-4360-abd7-d7f6c42d3f36","hostname":"api-appsflyer-5fbddc87f8-n6vp4","type":"filebeat","version":"7.4.0","ephemeral_id":"5dcfeba6-a8bf-460e-937e-67ccbb59a711"},"app_version":"3.0.2","language":"中文","input":{"type":"log"},"campaign_type":"organic","carrier":"CHINA TELECOM","ecs":{"version":"1.1.0"},"state":"BJ","postal_code":"100000","BasicData":"1,29,3,302,0,4de318f461e57002,0","conversion_type":"install","event_source":"SDK","appsflyer_id":"1702380351991-1046012330971770617","os_version":"10","event_value":"{\"uid\":\"2988393600\",\"publicProperty\":\"{\\\"userId\\\":\\\"2988393600\\\",\\\"deviceId\\\":\\\"1,29,3,302,0,4de318f461e57002,0\\\",\\\"languageCode\\\":\\\"zh-CN\\\",\\\"countryCode\\\":\\\"156\\\",\\\"env\\\":\\\"正式环境\\\",\\\"onlyid\\\":\\\"4de318f461e57002\\\",\\\"isv3\\\":\\\"旧V3\\\"}\",\"url\":\"https:\\/\\/wikidc.interface004.com\\/fxapp\\/fxcfg.json\"}","event_time_selected_timezone":"2024-01-29 04:01:42.364+0800","operator":"China Telecom","selected_timezone":"Asia/Shanghai","is_primary_attribution":true,"@version":"1","event_name":"af_appfly_updata_url","log":{"flags":["multiline"],"offset":1172304973,"file":{"path":"/var/app-appsflyer/filebeat/filebeat-2024-01-29.json"}}}`)
			producer.AsyncMessage("user_behavior_data_pipline", msg)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 500000; i++ {
			// TODO
			msg := sarama.StringEncoder(`{"@version":"1","BasicData":"999,0,4,999,0,2b29d85d93081fd520c027002858f899,0","code":"6172809985820","origin":"","isMinal":false,"name":"USDC","WikiChannel":"","project":0,"type":0,"@timestamp":"2024-02-13T09:26:39.435Z","ip":"103.169.64.57","isInternational":false,"UniqueDeviceId":"2b29d85d93081fd520c027002858f899","clickListType":4,"log":{"file":{"path":"/data/wikibitsearch.collect/info.log"},"offset":1470999,"flags":["multiline"]},"IpDeductionCity":"拉瓦尔品第","query":0,"log_date":"2024-02-13T09:26:39.000+0000","product":3,"channel":1,"agent":{"type":"filebeat","hostname":"api-wikibitsearch-bd8bd48df-54q26","ephemeral_id":"48bef397-d339-4cfa-8fc6-db007f292564","id":"fc2d3e96-41f9-4e73-ade4-762ce8afec23","version":"7.4.0"},"tags":["k8s_business","ACK-SG","beats_input_codec_json_applied"],"input":{"type":"log"},"business":"wikibitsearch.collect","logstash_id":"6690862b-94e2-41ce-bf91-a902e17a8a14","CreateTime":"2024-02-13 17:26:39","uid":"7722676852","host":{"mac":["00:16:3e:0a:9b:95"],"name":"api-wikibitsearch-bd8bd48df-54q26","architecture":"x86_64","hostname":"api-wikibitsearch-bd8bd48df-54q26","ip":["172.21.55.198","fe80::16:3e00:300a:9b95"],"os":{"version":"","name":"Alpine Linux","family":"","platform":"alpine","kernel":"5.10.134-13.al8.x86_64"},"containerized":false},"domain":"www.wikibit.com","clientId":"","language":"en","paltform":2,"keyword":"USDC","ecs":{"version":"1.1.0"},"country":"586","IpDeductionCountry":"巴基斯坦","searchOrigin":0,"clickTraderPosition":0}`)
			producer.AsyncMessage("wikifx-collect", msg)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 500000; i++ {
			// TODO
			msg := sarama.StringEncoder(`{
    "PageCode": "",
    "Platform": "999,0,2,999,0,5cc555dd04b7d0944c470038da7e3e1e,0",
    "CreateTime": "2024-03-07 10:58:18",
    "ecs": {
        "version": "1.1.0"
    },
    "agent": {
        "ephemeral_id": "ce14548d-11f0-4049-b51d-74532816ff7d",
        "version": "7.4.0",
        "hostname": "logscollection-business-2nbkc",
        "id": "03d8e238-c5d1-4d53-a031-20555e27bea6",
        "type": "filebeat"
    },
    "PageType": "",
    "IsInternational": false,
    "UserAgent": "",
    "Type": 1,
    "@version": "1",
    "ResultTotal": 4,
    "@timestamp": "2024-03-07T02:58:21.640Z",
    "LogDate": "",
    "SearchPosition": "",
    "SearchContent": "invinox",
    "PromoteSource": "",
    "Channel": "Web",
    "CountryCode": "826",
    "Product": "WikiFX",
    "latitude": 51.5245,
    "Source": "Website",
    "RegulatorCode": "",
    "BasicData": "999,0,2,999,0,5cc555dd04b7d0944c470038da7e3e1e,0",
    "logstash_id": "dec2051e-a9d2-4b37-8563-372166ef90a3",
    "log_date": "2024-03-07T02:58:18.000+0000",
    "Origin": "Website",
    "log": {
        "file": {
            "path": "/var/business/wikibitsearch.statistics/info.log"
        },
        "flags": [
            "multiline"
        ],
        "offset": 1401285
    },
    "SearchUrl": "",
    "business": "wikibitsearch.statistics",
    "Project": 1,
    "input": {
        "type": "log"
    },
    "FingerPrint": "",
    "IsMinal": false,
    "OverCordon": false,
    "City": "Tower Hamlets",
    "longitude": -0.0255,
    "Country": "英国",
    "IpAddress": "92.5.224.147",
    "LanguageCode": "en",
    "WikiChannel": "",
    "MaxMatch": 19.86,
    "host": {
        "mac": [
            "00:16:3e:04:3b:9f"
        ],
        "architecture": "x86_64",
        "containerized": false,
        "name": "logscollection-business-2nbkc",
        "hostname": "logscollection-business-2nbkc",
        "ip": [
            "172.21.42.164",
            "fe80::16:3e02:8c04:3b9f"
        ],
        "os": {
            "kernel": "5.10.134-13.al8.x86_64",
            "version": "",
            "name": "Alpine Linux",
            "family": "",
            "platform": "alpine"
        }
    },
    "SearchOrigin": 0,
    "ClientId": "",
    "UniqueDeviceId": "5cc555dd04b7d0944c470038da7e3e1e",
    "tags": [
        "k8s_business",
        "ACK-SG",
        "beats_input_codec_json_applied"
    ]
}`)
			producer.AsyncMessage("wikibit_search", msg)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 500000; i++ {
			msg := sarama.StringEncoder(`{"jumpStatisticsUrl":null,"Fluid":"253217ba-9635-4e82-9ed5-9d4c29fa702a","wikiChannel":"","CreatedDate":"0001-01-01T00:00:00.000Z","userid":"0","code":"1290453790","entry":"0","ipProvince":"布宜诺斯艾利斯","ver":"310","UtcTime":"2024-02-22T17:05:24.064Z","sandbox":"0","lang":"es","Status":"0","IsCharge":"0","platform":"3","ischinasystem":"False","type":"1","UtcDate":"2024-02-22","@timestamp":"2024-02-22T17:05:25.202Z","ip":"181.9.199.220","log":{"file":{"path":"/var/business/wikifxStatisitcs/20240223.log"},"offset":839231,"flags":["multiline"]},"ipCity":"布宜诺斯艾利斯","log_date":"2024-02-22T17:05:25.202Z","spots":"5","input":{"type":"log"},"tags":["k8s_business","ACK-SH","beats_input_codec_json_applied","_rubyexception"],"agent":{"type":"filebeat","hostname":"logscollection-business-4mrgv","ephemeral_id":"39170814-d054-4ca6-a058-b34e0cc160ce","id":"8b660606-0bf4-465c-88d7-4c9c43ef5c7d","version":"7.4.0"},"url":"","logstash_id":"4baa0735-00e3-44a4-b303-347db7a27c92","Remote_Addr":null,"business":"wikifxStatisitcs","host":{"mac":["00:16:3e:10:6b:e3"],"name":"logscollection-business-4mrgv","architecture":"x86_64","hostname":"logscollection-business-4mrgv","ip":["172.19.126.68","fe80::16:3e05:7110:6be3"],"os":{"version":"","name":"Alpine Linux","family":"","platform":"alpine","kernel":"5.10.134-13.al8.x86_64"},"containerized":false},"modal":"1","jumpStatisticsUrlParm":null,"area":"45030053","ipCountry":"阿根廷","equipinfo":null,"apptype":"1","ecs":{"version":"1.1.0"},"country":"032","equipId":"6f1da20c9a561414","@version":"1"}`)
			producer.AsyncMessage("wikifx-statistics", msg)
		}
		wg.Done()
	}()

	wg.Wait()
}

func TestNewProducer(t *testing.T) {
	producer, err := NewProducer(&ProducerConfig{
		Brokers: []string{"testkfk.fxeyeinterface.com:9092"},
		Version: "3.6.0",
	})
	assert.Nil(t, err)

	// 现有100个人安装
	// 50个iOS
	for i := 0; i < 50000; i++ {
		message := &AppFlyer{
			EventTime:   "2024-01-01 12:12:12",
			EventName:   "install",
			Platform:    "iOS",
			AppsflyerId: strconv.Itoa(i),
			AppName:     "Wiki global",
		}

		str, _ := json.Marshal(message)
		msg := sarama.StringEncoder(str)
		producer.AsyncMessage("user_behavior_data_pipline", msg)
	}

	// 50个Android
	for i := 0; i < 50; i++ {
		message := &AppFlyer{
			EventTime:   "2024-01-01 12:12:12",
			EventName:   "install",
			Platform:    "Android",
			AppsflyerId: strconv.Itoa(i),
			AppName:     "Wiki global",
		}

		str, _ := json.Marshal(message)
		msg := sarama.StringEncoder(str)
		producer.AsyncMessage("user_behavior_data_pipline", msg)
	}

	restartFunc := func(number, day int, os string) {
		eventTime, _ := time.Parse(time.DateTime, "2024-01-01 12:12:12")
		for i := 0; i < number; i++ {
			message := &AppFlyer{
				EventTime:   eventTime.AddDate(0, 0, day).Format(time.DateTime),
				EventName:   "af_restart",
				Platform:    os,
				AppsflyerId: strconv.Itoa(i),
				AppName:     "Wiki global",
			}

			str, _ := json.Marshal(message)
			msg := sarama.StringEncoder(str)
			producer.AsyncMessage("user_behavior_data_pipline", msg)
		}
	}

	// iOS中
	//次日有 40个打开
	restartFunc(40, 1, "iOS")
	//3日有35个打开
	restartFunc(35, 3, "iOS")
	//7日有20个打开
	restartFunc(20, 7, "iOS")
	//15日有15个打开
	restartFunc(15, 15, "iOS")
	//30日有10个打开
	restartFunc(10, 30, "iOS")

	// Android中
	//次日有 40个打开
	restartFunc(40, 1, "Android")
	//3日有35个打开
	restartFunc(35, 3, "Android")
	//7日有20个打开
	restartFunc(20, 7, "Android")
	//15日有15个打开
	restartFunc(15, 15, "Android")
	//30日有10个打开
	restartFunc(10, 30, "Android")

	producer.Close()
}

func TestNewProducer2(t *testing.T) {
	t.SkipNow()
	producer, err := NewProducer(&ProducerConfig{
		Brokers: []string{"testkfk.fxeyeinterface.com:9092"},
		Version: "3.6.0",
	})
	assert.Nil(t, err)

	// 现有100个人安装
	// 50个iOS
	for i := 0; i < 1; i++ {
		message := &AppFlyer{
			EventTime:   "2024-01-01 12:12:12",
			EventName:   "install",
			Platform:    "iOS",
			AppsflyerId: strconv.Itoa(i),
			AppName:     "Wiki global",
		}

		str, _ := json.Marshal(message)
		msg := sarama.StringEncoder(str)
		producer.AsyncMessage("user_behavior_data_pipline", msg)
	}

	restartFunc := func(number, day int, os string) {
		eventTime, _ := time.Parse(time.DateTime, "2024-01-01 12:12:12")
		for i := 0; i < number; i++ {
			message := &AppFlyer{
				EventTime:   eventTime.AddDate(0, 0, day).Format(time.DateTime),
				EventName:   "af_restart",
				Platform:    os,
				AppsflyerId: strconv.Itoa(i),
				AppName:     "Wiki global",
			}

			str, _ := json.Marshal(message)
			msg := sarama.StringEncoder(str)
			producer.AsyncMessage("user_behavior_data_pipline", msg)
		}
	}

	// iOS中
	//次日有 40个打开
	restartFunc(1, 30, "iOS")

	producer.Close()
}

func TestNewAWSProducer(t *testing.T) {
	t.SkipNow()
	producer, err := NewProducer(&ProducerConfig{
		Brokers: []string{
			"b-1-public.datawarehouse.4xla5g.c3.kafka.ap-southeast-1.amazonaws.com:9196",
			"b-2-public.datawarehouse.4xla5g.c3.kafka.ap-southeast-1.amazonaws.com:9196",
			"b-3-public.datawarehouse.4xla5g.c3.kafka.ap-southeast-1.amazonaws.com:9196",
		},
		Username:   "msk-data-warehouse",
		Password:   "Wikifx2024",
		EnableSASL: true,
		EnableTLS:  true,
		Version:    "3.5.1",
		Algorithm:  "sha512",
	})
	assert.Nil(t, err)

	value := []byte(`{"jumpStatisticsUrl":null,"Fluid":"253217ba-9635-4e82-9ed5-9d4c29fa702a","wikiChannel":"","CreatedDate":"0001-01-01T00:00:00.000Z","userid":"0","code":"1290453790","entry":"0","ipProvince":"布宜诺斯艾利斯","ver":"310","UtcTime":"2024-02-22T17:05:24.064Z","sandbox":"0","lang":"es","Status":"0","IsCharge":"0","platform":"3","ischinasystem":"False","type":"1","UtcDate":"2024-02-22","@timestamp":"2024-02-22T17:05:25.202Z","ip":"181.9.199.220","log":{"file":{"path":"/var/business/wikifxStatisitcs/20240223.log"},"offset":839231,"flags":["multiline"]},"ipCity":"布宜诺斯艾利斯","log_date":"2024-02-22T17:05:25.202Z","spots":"5","input":{"type":"log"},"tags":["k8s_business","ACK-SH","beats_input_codec_json_applied","_rubyexception"],"agent":{"type":"filebeat","hostname":"logscollection-business-4mrgv","ephemeral_id":"39170814-d054-4ca6-a058-b34e0cc160ce","id":"8b660606-0bf4-465c-88d7-4c9c43ef5c7d","version":"7.4.0"},"url":"","logstash_id":"4baa0735-00e3-44a4-b303-347db7a27c92","Remote_Addr":null,"business":"wikifxStatisitcs","host":{"mac":["00:16:3e:10:6b:e3"],"name":"logscollection-business-4mrgv","architecture":"x86_64","hostname":"logscollection-business-4mrgv","ip":["172.19.126.68","fe80::16:3e05:7110:6be3"],"os":{"version":"","name":"Alpine Linux","family":"","platform":"alpine","kernel":"5.10.134-13.al8.x86_64"},"containerized":false},"modal":"1","jumpStatisticsUrlParm":null,"area":"45030053","ipCountry":"阿根廷","equipinfo":null,"apptype":"1","ecs":{"version":"1.1.0"},"country":"032","equipId":"6f1da20c9a561414","@version":"1"}`)
	for i := 0; i < 1000; i++ {
		fmt.Println(i)
		producer.AsyncMessage("mytopic", sarama.ByteEncoder(value))
	}
	producer.Close()
}
