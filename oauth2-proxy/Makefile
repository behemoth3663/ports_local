PORTNAME=	oauth2-proxy
PORTVERSION=	7.15.5
DISTVERSIONPREFIX=	v
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
GO_LDFLAGS+=	-w
GO_LDFLAGS+=	-X version.VERSION=${PORTVERSION}
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

M2T_OFFLINE=	true
.export	M2T_OFFLINE

post-install: .PHONY
	${MKDIR} ${STAGEDIR}${LOG_DIR}
	${MKDIR} ${STAGEDIR}${RUN_DIR}
	${INSTALL_DATA} ${WRKSRC}/contrib/${PORTNAME}.cfg.example ${STAGEDIR}${PREFIX}/etc/${PORTNAME}.cfg.sample

_OPTIONS_extract=	750:post-post-extract
post-post-extract: .PHONY
	@${MKDIR} ${WRKSRC}/vendor
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@${RM} ${WRKSRC}/vendor/github.com/googleapis/gax-go/v2/telemetry.go

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

post-deps-all: .PHONY
	${TEST} ! -d '${PATCHDIR}_' || ${MV} ${PATCHDIR}_ ${PATCHDIR}

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=go-yaml GHT_PROJ=yaml GHT_PATH=go.yaml.in/yaml/v3 )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:blank:]]*#[[:blank:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:blank:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-google:go-genproto:6ceb2ff114de:google_go_genproto/vendor/google.golang.org/genproto/googleapis/rpc
#+google:go-genproto:6ceb2ff114de:google_go_genproto/vendor/google.golang.org/genproto
GH_TUPLE:=	${GH_TUPLE:S|/googleapis/rpc$||1}

#-googleapis:gax-go:v2.11.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go/v2
#+googleapis:gax-go:v2.11.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go
GH_TUPLE:=	${GH_TUPLE:S|/gax-go/v2$|/gax-go|1}

#-googleapis:google-cloud-go:v0.23.3:googleapis_google_cloud_go_auth/vendor/cloud.google.com/go/auth
#+googleapis:google-cloud-go:auth/v0.23.3:googleapis_google_cloud_go_auth/vendor/cloud.google.com/go
GH_TUPLE:=	${GH_TUPLE:C,:([^:]+:[^:/]+/.*/cloud.google.com/go)/(auth),:\2/\1,1}

.sinclude "${.CURDIR}/generated.deps.mk"

#-go-yaml:yaml:v3.0.5:go_yaml_yaml/vendor/go.yaml.in/yaml/v3
#+yaml:go-yaml:v3.0.5:go_yaml_yaml/vendor/go.yaml.in/yaml/v3
GH_TUPLE:=	${GH_TUPLE:S|go-yaml:yaml:|yaml:go-yaml:|1}

.include <bsd.port.mk>
