PORTNAME=	oauth2-proxy
PORTVERSION=	7.15.4
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local www
PKGNAMESUFFIX=	-local

MAINTAINER=	freebsd@rheinwolf.de
COMMENT=	OAuth 2.0 proxy server
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

USES=		go:modules

USE_RC_SUBR=	oauth2-proxy

USE_GITHUB=	yes

USER=		www
GROUP=		www

VAR_DIR?=	/var
LOG_DIR?=	${VAR_DIR}/log/${PORTNAME}
RUN_DIR?=	${VAR_DIR}/run/${PORTNAME}

SUB_LIST+=	GROUP=${GROUP} \
		LOG_DIR=${LOG_DIR} \
		PORTNAME=${PORTNAME} \
		PREFIX=${PREFIX} \
		RUN_DIR=${RUN_DIR} \
		USER=${USER}

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X version.VERSION=${PORTVERSION}"
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

M2T_OFFLINE=	true
.export	M2T_OFFLINE

post-install: .PHONY
	${MKDIR} ${STAGEDIR}${LOG_DIR}
	${MKDIR} ${STAGEDIR}${RUN_DIR}
	${INSTALL_DATA} ${WRKSRC}/contrib/${PORTNAME}.cfg.example ${STAGEDIR}${PREFIX}/etc/${PORTNAME}.cfg.sample

_OPTIONS_extract=	750:post-post-extract
post-post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@${RM} ${WRKSRC}/vendor/github.com/googleapis/gax-go/v2/telemetry.go
	@${MKDIR} ${WRKSRC}/vendor/go.yaml.in/yaml
	@${RLN} ${WRKSRC}/vendor/gopkg.in/yaml.v3 ${WRKSRC}/vendor/go.yaml.in/yaml/v3

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

post-deps-all: .PHONY
	${TEST} ! -d '${PATCHDIR}_' || ${MV} ${PATCHDIR}_ ${PATCHDIR}

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-google:go-genproto:6ceb2ff114de:google_go_genproto/vendor/google.golang.org/genproto/googleapis/rpc
#+google:go-genproto:6ceb2ff114de:google_go_genproto/vendor/google.golang.org/genproto
GH_TUPLE:=	${GH_TUPLE:S|/googleapis/rpc$||}

#-googleapis:gax-go:v2.11.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go/v2
#+googleapis:gax-go:v2.11.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go
GH_TUPLE:=	${GH_TUPLE:S|/gax-go/v2$|/gax-go|1}

#-googleapis:google-cloud-go:v0.2.5:googleapis_google_cloud_go/vendor/cloud.google.com/go/auth/oauth2adapt
#+googleapis:google-cloud-go:auth/oauth2adapt/v0.2.5:googleapis_google_cloud_go/vendor/cloud.google.com/go
#-googleapis:google-cloud-go:v0.5.2:googleapis_google_cloud_go_1/vendor/cloud.google.com/go/compute/metadata
#+googleapis:google-cloud-go:compute/metadata/v0.5.2:googleapis_google_cloud_go_1/vendor/cloud.google.com/go
GH_TUPLE:=	${GH_TUPLE:C|^(googleapis:google-cloud-go:)(.*/go)/(.*)$|\1\3/\2|}

.include <bsd.port.mk>
