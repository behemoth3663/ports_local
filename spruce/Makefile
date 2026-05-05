PORTNAME=	spruce
DISTVERSIONPREFIX=	v
DISTVERSION=	1.35.1
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	YAML & JSON merging tool
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

USES=	go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	geofffranks

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X 'main.Version=${DISTVERSIONFULL}'"

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

# The vendor directory is inside geofffranks-spruce-v1.35.1_GH0.tar.gz, so there's no need to copy modules.txt for now.
#-aws:aws-sdk-go-v2:v1.41.6:aws_aws_sdk_go_v2_secretsmanager/vendor/github.com/aws/aws-sdk-go-v2
#-aws:aws-sdk-go-v2:v1.41.7:aws_aws_sdk_go_v2
#+aws:aws-sdk-go-v2:v1.41.7:aws_aws_sdk_go_v2/vendor/github.com/aws/aws-sdk-go-v2
#GH_TUPLE:=	${GH_TUPLE:Naws*aws_aws_sdk_go_v2_*}
#GH_TUPLE:=	${GH_TUPLE:C|:aws_aws_sdk_go_v2|&/vendor/github.com/aws/aws-sdk-go-v2|}

#-open-telemetry:opentelemetry-go-contrib:v0.61.0:open_telemetry_opentelemetry_go_contrib/vendor/go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp
#+open-telemetry:opentelemetry-go-contrib:instrumentation/net/http/otelhttp/v0.61.0:open_telemetry_opentelemetry_go_contrib/vendor/go.opentelemetry.io/contrib
#-open-telemetry:opentelemetry-go-instrumentation:v1.2.1:open_telemetry_opentelemetry_go_instrumentation/vendor/go.opentelemetry.io/auto/sdk
#+open-telemetry:opentelemetry-go-instrumentation:sdk/v1.2.1:open_telemetry_opentelemetry_go_instrumentation/vendor/go.opentelemetry.io/auto
GH_TUPLE:=	${GH_TUPLE:C|^(open-telemetry:[^:]+:)([^:]+)(.*/vendor/go.opentelemetry.io/[^/]+)/(.*)|\1\4/\2\3|}

#-starkandwayne:goutils:896b8a6904be:starkandwayne_goutils/vendor/github.com/starkandwayne/goutils
#+egen:goutils:896b8a6904be:starkandwayne_goutils/vendor/github.com/starkandwayne/goutils
GH_TUPLE:=	${GH_TUPLE:C|^starkandwayne:|egen:|}

.include <bsd.port.mk>
