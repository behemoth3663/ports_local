PORTNAME=	cloud-nuke
PORTVERSION=	0.52.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils
PKGNAMESUFFIX=	-local

MAINTAINER=	none@freebsd.org
COMMENT=	CLI tool to delete all resources in an AWS account
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE.txt

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	gruntwork-io

PLIST_FILES=	bin/${PORTNAME}

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w -X main.VERSION=${DISTVERSIONFULL}"
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

M2T_OFFLINE=	true
.export	M2T_OFFLINE

_OPTIONS_extract=	600:pre-post-extract
#Resolve the error:
#mv: rename /usr/obj/usr/ports/local/cloud-nuke/work/gruntwork-cli-0.7.0 to /usr/obj/usr/ports/local/cloud-nuke/work/cloud-nuke-0.29.4/vendor/github.com/gruntwork-io/gruntwork-cli: No such file or directory
#/usr/obj/usr/ports/local/cloud-nuke/work/gruntwork-cli-0.7.0 -> /usr/obj/usr/ports/local/cloud-nuke/work/go-commons-0.7.0
#/usr/obj/usr/ports/local/cloud-nuke/work/gruntwork-cli-0.16.1 -> /usr/obj/usr/ports/local/cloud-nuke/work/go-commons-0.16.1
pre-post-extract: .PHONY .SILENT
	${TEST} -z ${WRKDIR} || ${FIND} ${WRKDIR} -maxdepth 1 -type d -name 'go-commons-*' | while read src; do \
		${LN} -s "$${src}" "$${src%%go-commons-*}gruntwork-cli-$${src##*go-commons-}"; \
	done

AWS_SDK_GO_V2+=aws/protocol/eventstream
AWS_SDK_GO_V2+=config
AWS_SDK_GO_V2+=credentials
AWS_SDK_GO_V2+=feature/ec2/imds
AWS_SDK_GO_V2+=internal/configsources
AWS_SDK_GO_V2+=internal/ini
AWS_SDK_GO_V2+=service/accessanalyzer
AWS_SDK_GO_V2+=service/acm
AWS_SDK_GO_V2+=service/acmpca
AWS_SDK_GO_V2+=service/amp
AWS_SDK_GO_V2+=service/apigateway
AWS_SDK_GO_V2+=service/apigatewayv2
AWS_SDK_GO_V2+=service/apprunner
AWS_SDK_GO_V2+=service/autoscaling
AWS_SDK_GO_V2+=service/backup
AWS_SDK_GO_V2+=service/cloudformation
AWS_SDK_GO_V2+=service/cloudfront
AWS_SDK_GO_V2+=service/cloudtrail
AWS_SDK_GO_V2+=service/cloudwatch
AWS_SDK_GO_V2+=service/cloudwatchlogs
AWS_SDK_GO_V2+=service/codedeploy
AWS_SDK_GO_V2+=service/configservice
AWS_SDK_GO_V2+=service/datapipeline
AWS_SDK_GO_V2+=service/datasync
AWS_SDK_GO_V2+=service/dynamodb
AWS_SDK_GO_V2+=service/ec2
AWS_SDK_GO_V2+=service/ecr
AWS_SDK_GO_V2+=service/ecs
AWS_SDK_GO_V2+=service/efs
AWS_SDK_GO_V2+=service/eks
AWS_SDK_GO_V2+=service/elasticache
AWS_SDK_GO_V2+=service/elasticbeanstalk
AWS_SDK_GO_V2+=service/elasticloadbalancing
AWS_SDK_GO_V2+=service/elasticloadbalancingv2
AWS_SDK_GO_V2+=service/eventbridge
AWS_SDK_GO_V2+=service/firehose
AWS_SDK_GO_V2+=service/grafana
AWS_SDK_GO_V2+=service/guardduty
AWS_SDK_GO_V2+=service/iam
AWS_SDK_GO_V2+=service/internal/accept-encoding
AWS_SDK_GO_V2+=service/internal/checksum
AWS_SDK_GO_V2+=service/internal/endpoint-discovery
AWS_SDK_GO_V2+=service/internal/presigned-url
AWS_SDK_GO_V2+=service/internal/s3shared
AWS_SDK_GO_V2+=service/kafka
AWS_SDK_GO_V2+=service/kinesis
AWS_SDK_GO_V2+=service/kms
AWS_SDK_GO_V2+=service/lambda
AWS_SDK_GO_V2+=service/macie2
AWS_SDK_GO_V2+=service/mq
AWS_SDK_GO_V2+=service/networkfirewall
AWS_SDK_GO_V2+=service/opensearch
AWS_SDK_GO_V2+=service/ram
AWS_SDK_GO_V2+=service/rds
AWS_SDK_GO_V2+=service/redshift
AWS_SDK_GO_V2+=service/route53
AWS_SDK_GO_V2+=service/s3
AWS_SDK_GO_V2+=service/s3control
AWS_SDK_GO_V2+=service/sagemaker
AWS_SDK_GO_V2+=service/scheduler
AWS_SDK_GO_V2+=service/secretsmanager
AWS_SDK_GO_V2+=service/securityhub
AWS_SDK_GO_V2+=service/servicediscovery
AWS_SDK_GO_V2+=service/ses
AWS_SDK_GO_V2+=service/sns
AWS_SDK_GO_V2+=service/sqs
AWS_SDK_GO_V2+=service/ssm
AWS_SDK_GO_V2+=service/sso
AWS_SDK_GO_V2+=service/ssooidc
AWS_SDK_GO_V2+=service/sts
AWS_SDK_GO_V2+=service/vpclattice

post-extract: .PHONY .SILENT
	${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	${RM} -r ${WRKSRC}/vendor/google.golang.org/grpc/stats/opentelemetry
	${MKDIR} ${WRKSRC}/vendor/github.com/aws/aws-sdk-go-v2/internal
	${TEST} -z ${WRKSRC_aws_aws_sdk_go_v2_internal} || ${RLN} ${WRKSRC_aws_aws_sdk_go_v2_internal}/internal ${WRKSRC}/vendor/github.com/aws/aws-sdk-go-v2
.for _i in ${AWS_SDK_GO_V2}
	${RM} -r ${WRKSRC}/vendor/github.com/aws/aws-sdk-go-v2/${_i}
	${TEST} -z ${WRKSRC_aws_aws_sdk_go_v2_${_i:T:S/-/_/}} || ${RLN} ${WRKSRC_aws_aws_sdk_go_v2_${_i:T:S/-/_/}}/${_i} ${WRKSRC}/vendor/github.com/aws/aws-sdk-go-v2/${_i}
.endfor
	${RM} -r ${WRKSRC}/vendor/github.com/aws/aws-sdk-go-v2/internal/endpoints/v2
	${TEST} -z ${WRKSRC_aws_aws_sdk_go_v2_v2} || ${RLN} ${WRKSRC_aws_aws_sdk_go_v2_v2}/internal/endpoints/v2 ${WRKSRC}/vendor/github.com/aws/aws-sdk-go-v2/internal/endpoints

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=atomicgo GHT_PROJ=cursor GHT_PATH=atomicgo.dev/cursor )
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=atomicgo GHT_PROJ=keyboard GHT_PATH=atomicgo.dev/keyboard )
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=google GHT_PROJ=cel-spec GHT_PATH=cel.dev/expr )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:blank:]]*#[[:blank:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:blank:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

post-deps-all: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${MV} ${PATCHDIR}_ ${PATCHDIR} || :

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-GoogleCloudPlatform:opentelemetry-operations-go:v0.50.0:googlecloudplatform_opentelemetry_operations_go_metric/vendor/github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric
#-GoogleCloudPlatform:opentelemetry-operations-go:v0.50.0:googlecloudplatform_opentelemetry_operations_go_resourcemapping/vendor/github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/resourcemapping
#-GoogleCloudPlatform:opentelemetry-operations-go:v1.30.0:googlecloudplatform_opentelemetry_operations_go_gcp/vendor/github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp
GH_TUPLE:=	${GH_TUPLE:NGoogleCloudPlatform\:opentelemetry-operations-go\:*}

#-aws-sdk-go-v2:v2.7.21:aws_aws_sdk_go_v2_v2/vendor/github.com/aws/aws-sdk-go-v2/internal/endpoints/v2
#+aws-sdk-go-v2:internal/endpoints/v2.7.21:aws_aws_sdk_go_v2_v2
GH_TUPLE:=	${GH_TUPLE:C,^([^:]+:[^:]+:)([^:]+:[^/]+)/vendor/github.com/aws/aws-sdk-go-v2/(internal/endpoints)/v2$,\1\3/\2,}

#-aws:aws-sdk-go-v2:v1.55.8:aws_aws_sdk_go_v2_securityhub/vendor/github.com/aws/aws-sdk-go-v2/service/securityhub
#+aws:aws-sdk-go-v2:service/securityhub/v1.55.8:aws_aws_sdk_go_v2_securityhub
GH_TUPLE:=	${GH_TUPLE:C,^([^:]+:[^:]+:)([^:]+:[^/]+)/vendor/github.com/aws/aws-sdk-go-v2/([^[:blank:]]+)$,\1\3/\2,}

# googleapis:google-cloud-go:v0.112.1:googleapis_google_cloud_go/vendor/cloud.google.com/go
#-googleapis:google-cloud-go:v1.1.6:googleapis_google_cloud_go_iam/vendor/cloud.google.com/go/iam
#-googleapis:google-cloud-go:v1.39.1:googleapis_google_cloud_go_storage/vendor/cloud.google.com/go/storage
GH_TUPLE:=	${GH_TUPLE:Ngoogleapis\:google-cloud-go\:*/cloud.google.com/go/*}

#-envoyproxy:go-control-plane:v1.32.4:envoyproxy_go_control_plane_envoy/vendor/github.com/envoyproxy/go-control-plane/envoy
#+envoyproxy:go-control-plane:envoy/v1.32.4:envoyproxy_go_control_plane_envoy/vendor/github.com/envoyproxy/go-control-plane
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+/vendor/github.com/envoyproxy/go-control-plane)/(envoy)|\1\3/\2|}

#-googleapis:gax-go:v2.13.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go/v2
#+googleapis:gax-go:v2.13.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go
GH_TUPLE:=	${GH_TUPLE:S|/gax-go/v2|/gax-go|1}

#-cncf:xds:ee656c7534f5:cncf_xds_go/vendor/github.com/cncf/xds/go
#+cncf:xds:ee656c7534f5:cncf_xds_go/vendor/github.com/cncf/xds
GH_TUPLE:=	${GH_TUPLE:S|/xds/go|/xds|1}

.sinclude "${.CURDIR}/generated.deps.mk"
.include <bsd.port.mk>
