PORTNAME=	gosec
PORTVERSION=	2.29.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local devel security
PKGNAMESUFFIX=	-local

MAINTAINER=	none@freebsd.org
COMMENT=	Golang Security Checker
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE.txt

BUILD_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	securego

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

COMMIT_CMD=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONFULL} | cut -c1-7

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X 'main.GitTag=$$(${COMMIT_CMD})' \
	-X 'main.Version=${PORTVERSION}' \
	-X 'main.BuildDate=${%F %T:L:gmtime}'"
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

M2T_OFFLINE=	true
.export	M2T_OFFLINE

post-extract: .PHONY .SILENT
	${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	${RM} -r ${WRKSRC}/vendor/cloud.google.com/go/auth
	${TEST} -z ${WRKSRC_googleapis_google_cloud_go_auth} || ${RLN} ${WRKSRC_googleapis_google_cloud_go_auth}/auth ${WRKSRC}/vendor/cloud.google.com/go
	${RM} -r ${WRKSRC}/vendor/cloud.google.com/go/compute/metadata
	${TEST} -z ${WRKSRC_googleapis_google_cloud_go_compute_metadata} || ${RLN} ${WRKSRC_googleapis_google_cloud_go_compute_metadata}/compute/metadata ${WRKSRC}/vendor/cloud.google.com/go/compute

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

post-deps-all: .PHONY
	${TEST} ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=googleapis GHT_PROJ=google-cloud-go GHT_PATH=cloud.google.com/go/compute/metadata )
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=googleapis GHT_PROJ=go-genai GHT_PATH=google.golang.org/genai )
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=yaml GHT_PROJ=go-yaml GHT_PATH=go.yaml.in/yaml/v3 )
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=yaml GHT_PROJ=go-yaml GHT_PATH=go.yaml.in/yaml/v4 )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:blank:]]*#[[:blank:]]*::([-v0-9.a-frc]+):.*/(vendor/${GHT_PATH})[[:blank:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-google:go-genproto:08b0e4226688:google_go_genproto_rpc/vendor/google.golang.org/genproto/googleapis/rpc
#+google:go-genproto:08b0e4226688:google_go_genproto_rpc/vendor/google.golang.org/genproto
GH_TUPLE:=	${GH_TUPLE:S|/googleapis/rpc||1}

#-googleapis:gax-go:v2.14.1:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go/v2
#+googleapis:gax-go:v2.14.1:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go
GH_TUPLE:=	${GH_TUPLE:S|/gax-go/v2|/gax-go|1}

#-googleapis:google-cloud-go:v0.23.2:googleapis_google_cloud_go_auth/vendor/cloud.google.com/go/auth
#+googleapis:google-cloud-go:v0.23.2:googleapis_google_cloud_go_auth
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/cloud.google.com/go/(auth)$|\1\3/\2|}

#-standard-webhooks:standard-webhooks:v0.0.1:standard_webhooks_standard_webhooks_libraries/vendor/github.com/standard-webhooks/standard-webhooks/libraries
#+standard-webhooks:standard-webhooks:libraries/v0.0.1:standard_webhooks_standard_webhooks_libraries/vendor/github.com/standard-webhooks/standard-webhooks
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)(.*/standard-webhooks/standard-webhooks)/(libraries)$|\1\3/\2|1}

.sinclude "${.CURDIR}/generated.deps.mk"

GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/cloud.google.com/go/(compute)/(metadata)$|\1\3/\4/\2_\3_\4|}

#-yaml:go-yaml:v3.0.5:yaml_go_yaml/vendor/go.yaml.in/yaml/v3
#+yaml:go-yaml:v3.0.5:yaml_go_yaml/vendor/go.yaml.in/yaml/v3
#-yaml:go-yaml:v4.0.0-rc.6:yaml_go_yaml/vendor/go.yaml.in/yaml/v4
#+yaml:go-yaml:v4.0.0-rc.6:yaml_go_yaml/vendor/go.yaml.in/yaml/v4
GH_TUPLE:=	${GH_TUPLE:C|(/vendor/go.yaml.in/yaml/)(.*)|_\2\1\2|}

.include <bsd.port.mk>
