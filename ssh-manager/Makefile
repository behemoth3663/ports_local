PORTNAME=	ssh-manager
DISTVERSIONPREFIX=	v
DISTVERSION=	1.126.0
CATEGORIES=	local security

MAINTAINER=	none@freebsd.org
COMMENT=	SSH Key Manager for 1Password, Bitwarden and AWS S3
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	omegion

PLIST_FILES=	bin/${PORTNAME}

CGO_ENABLED=	0
GO_LDFLAGS+=	-w
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

M2T_OFFLINE=	true
.export	M2T_OFFLINE

post-patch: .PHONY
	${REINPLACE_CMD} \
		's/log\.Infoln(version)/version.Version = "${PORTVERSION}"; version.Time = "${SOURCE_DATE_EPOCH_CMD:sh}"; &/' \
	${WRKSRC}/cmd/version.go

post-extract: .PHONY
	${RM} -r ${WRKSRC}/vendor/github.com/census-instrumentation
	${RM} -r ${WRKSRC}/vendor/go.opencensus.io
	${RM} ${WRKSRC}/vendor/github.com/envoyproxy/go-control-plane/envoy/config/trace/v3/opencensus.pb.*go
	${RM} ${WRKSRC}/vendor/google.golang.org/api/transport/http/internal/propagation/http.go
	@${RM} -r ${WRKSRC}/vendor/cloud.google.com/go/storage
	@${TEST} -z ${WRKSRC_googleapis_google_cloud_go_storage} || ${RLN} ${WRKSRC_googleapis_google_cloud_go_storage}/storage ${WRKSRC}/vendor/cloud.google.com/go/storage

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

GH_TUPLE:=	${GH_TUPLE:Naws*aws_aws_sdk_go_v2_*/vendor/github.com/aws/aws-sdk-go-v2/*}

#-census-instrumentation:opencensus-go:v0.23.0:census_instrumentation_opencensus_go/vendor/go.opencensus.io
#-census-instrumentation:opencensus-proto:v0.3.0:census_instrumentation_opencensus_proto/vendor/github.com/census-instrumentation/opencensus-proto
GH_TUPLE:=	${GH_TUPLE:Ncensus-instrumentation\:*}

#-cespare:xxhash:v2.1.2:cespare_xxhash_v2/vendor/github.com/cespare/xxhash/v2
GH_TUPLE:=	${GH_TUPLE:Ncespare\:xxhash\:*}

#-cncf:udpa:04548b0d99d4:cncf_udpa_go/vendor/github.com/cncf/udpa/go
#+cncf:udpa:04548b0d99d4:cncf_udpa_go/vendor/github.com/cncf/udpa
GH_TUPLE:=	${GH_TUPLE:S|/udpa/go|/udpa|1}

#-cncf:xds:d92e9ce0af51:cncf_xds_go/vendor/github.com/cncf/xds/go
#+cncf:xds:d92e9ce0af51:cncf_xds_go/vendor/github.com/cncf/xds
GH_TUPLE:=	${GH_TUPLE:S|/xds/go|/xds|1}

#-googleapis:gax-go:v2.13.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go/v2
#+googleapis:gax-go:v2.13.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go
GH_TUPLE:=	${GH_TUPLE:S|/gax-go/v2|/gax-go|1}

#-googleapis:google-cloud-go:v1.18.2:googleapis_google_cloud_go_storage/vendor/cloud.google.com/go/storage
#+googleapis:google-cloud-go:storage/v1.18.2:googleapis_google_cloud_go_storage
GH_TUPLE:=	${GH_TUPLE:C,:([^:]+:[^:/]+)/.*/cloud.google.com/go/(storage),:\2/\1,1}

.include <bsd.port.pre.mk>

SOURCE_DATE_EPOCH_CMD=	date -ur \
			$$(${GREP} --max-count=1 TIMESTAMP ${DISTINFO_FILE} | ${CUT} -d' ' -f3) \
			'+%FT%TZ'

.include <bsd.port.post.mk>
