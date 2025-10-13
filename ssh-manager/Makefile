PORTNAME=	ssh-manager
DISTVERSIONPREFIX=	v
DISTVERSION=	1.75.0
CATEGORIES=	local security

MAINTAINER=	none@freebsd.org
COMMENT=	SSH Key Manager for 1Password, Bitwarden and AWS S3
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	omegion

PLIST_FILES=	bin/${PORTNAME}

# looks like -buildid=${PORTVERSION} is not working
GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static"

M2T_OFFLINE=	true
.export	M2T_OFFLINE

post-patch: .PHONY
	${REINPLACE_CMD} \
		's/log\.Infoln(version)/version.Version = "${PORTVERSION}"; version.Time = "${%FT%TZ:L:gmtime}"; &/' \
	${WRKSRC}/cmd/version.go

post-extract: .PHONY
	@${RM} -r ${WRKSRC}/vendor/cloud.google.com/go/storage
	@test -z ${WRKSRC_googleapis_google_cloud_go_storage} || ${RLN} ${WRKSRC_googleapis_google_cloud_go_storage}/storage ${WRKSRC}/vendor/cloud.google.com/go/storage

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

GH_TUPLE:=	${GH_TUPLE:Naws*aws_aws_sdk_go_v2_*/vendor/github.com/aws/aws-sdk-go-v2/*}

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

#-googleapis:google-cloud-go:        v1.18.2:googleapis_google_cloud_go_storage/vendor/cloud.google.com/go/storage
#+googleapis:google-cloud-go:storage/v1.18.2:googleapis_google_cloud_go_storage
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/cloud.google.com/go/(storage)|\1\3/\2|}

.include <bsd.port.mk>
