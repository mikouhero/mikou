module mikou

go 1.17

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fsnotify/fsnotify v1.5.1
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-gonic/gin v1.7.4
	github.com/go-playground/universal-translator v0.18.0
	github.com/go-playground/validator/v10 v10.9.0
	github.com/jinzhu/gorm v1.9.12
	github.com/robfig/cron/v3 v3.0.0
	github.com/spf13/viper v1.4.0
	go.uber.org/zap v1.19.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gorm.io/driver/mysql v1.2.1
	gorm.io/gorm v1.22.4
	gorm.io/plugin/dbresolver v1.1.0
)

require (
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/mitchellh/mapstructure v1.4.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/aa v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/aai v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/acp v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/advisor v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/af v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/afc v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ame v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ams v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/antiddos v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/apcas v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ape v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/api v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/apigateway v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/apm v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/as v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/asr v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/asw v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ba v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/batch v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bda v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bi v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bizlive v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bm v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bma v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bmeip v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bmlb v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bmvpc v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bri v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bsca v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/btoe v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cam v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/captcha v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/car v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/casb v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cat v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cbs v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ccc v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdc v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cds v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cfg v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cfs v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cfw v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/chdfs v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ciam v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cii v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cim v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cis v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ckafka v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/clb v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cloudaudit v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cloudhsm v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cls v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cme v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cmq v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cms v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cpdp v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cr v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cwp v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cws v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cynosdb v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dasb v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dataintegration v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dayu v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dbbrain v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dbdc v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dc v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dcdb v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dlc v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/domain v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/drm v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ds v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dtf v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dts v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/eb v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ecc v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ecdn v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ecm v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/eiam v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/eis v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/emr v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/es v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/facefusion v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/faceid v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/fmu v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ft v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gaap v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gme v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gpm v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gs v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gse v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/habo v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/hcm v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iai v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ic v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/icr v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ie v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iecp v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iir v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ims v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iot v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iotcloud v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iotexplorer v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iottid v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iotvideo v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iotvideoindustry v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/irp v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ivld v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/kms v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/live v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lowcode v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lp v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mariadb v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/market v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/memcached v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mgobe v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mmps v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mna v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mongodb v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mps v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mrs v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ms v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/msp v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mvj v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/nlp v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/npp v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/oceanus v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/organization v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/partners v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/pds v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/postgres v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/privatedns v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/pts v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/rce v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/redis v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/region v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/rkp v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/rp v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/rum v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/scf v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ses v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/smh v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/smpn v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/soe v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/solar v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sqlserver v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssa v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssl v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sslpod v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssm v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/taf v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tag v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tan v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tat v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tav v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tbaas v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tbm v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tbp v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcaplusdb v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcb v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcbr v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcex v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tci v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcm v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcr v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcss v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tdcpg v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tdid v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tdmq v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tem v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/teo v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/thpc v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tia v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tic v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ticm v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tics v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tiems v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tiia v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tione v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tiw v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tke v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tkgdq v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tms v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tmt v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/trp v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/trtc v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tse v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tsf v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tsw v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tts v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ump v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vm v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vms v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vod v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vpc v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/waf v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/wav v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/wedata v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/wss v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/yinsuda v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/youmall v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/yunjing v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/yunsou v1.0.498 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/zj v1.0.498 // indirect
	github.com/ugorji/go/codec v1.2.6 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
