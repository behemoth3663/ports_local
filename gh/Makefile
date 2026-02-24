PORTNAME=	gh
PORTVERSION=	2.87.3
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local devel
PKGNAMESUFFIX?=	-local

MAINTAINER=	lcook@FreeBSD.org
COMMENT=	GitHub's official command line tool
WWW=		https://cli.github.com

LICENSE=	MIT

RUN_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	cli
GH_PROJECT=	${GH_ACCOUNT}

GO_TARGET=	./cmd/${PORTNAME}

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/internal/build.Version=${PORTVERSION} \
	-X github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/internal/build.Date=${%FT%TZ:L:gmtime}"
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

_DOCS=		README.md docs/[!README]*.md
PORTDOCS=	${_DOCS:T}

OPTIONS_DEFINE=		DOCS MANPAGES
OPTIONS_DEFAULT=	MANPAGES
OPTIONS_SUB=		yes

M2T_OFFLINE=	true
.export	M2T_OFFLINE

do-build-MANPAGES-on:
	(cd ${WRKSRC} && \
		${SETENV} ${MAKE_ENV} ${GO_ENV} ${GO_CMD} run ${GO_BUILDFLAGS} \
		./cmd/gen-docs --man-page --doc-path docs/man 2>/dev/null)

post-install-DOCS-on:
	@${MKDIR} ${STAGEDIR}${DOCSDIR}
	(cd ${WRKSRC} && \
		${INSTALL_DATA} ${_DOCS} \
		${STAGEDIR}${DOCSDIR})

post-install-MANPAGES-on:
	${INSTALL_MAN} ${WRKSRC}/docs/man/*.1 \
		${STAGEDIR}${PREFIX}/share/man/man1

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=imdario GHT_PROJ=mergo GHT_PATH=dario.cat/mergo )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:space:]]*#[[:space:]]*::(v[0-9.]+):.*/(vendor/${GHT_PATH})[[:space:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/}_${GHT_PROJ:tl:S/-/_/}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

post-deps-all: .PHONY
	${TEST} ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@${RM} ${WRKSRC}/vendor/github.com/go-openapi/runtime/client/opentelemetry.go
	@${MKDIR} ${WRKSRC}/vendor/google.golang.org/genproto/googleapis
	@${TEST} -z ${WRKSRC_google_go_genproto_api} || ${RLN} ${WRKSRC_google_go_genproto_api}/googleapis/api ${WRKSRC}/vendor/google.golang.org/genproto/googleapis
	@${TEST} -z ${WRKSRC_google_go_genproto_rpc} || ${RLN} ${WRKSRC_google_go_genproto_rpc}/googleapis/rpc ${WRKSRC}/vendor/google.golang.org/genproto/googleapis
	@${RM} -r ${WRKSRC}/vendor/github.com/charmbracelet/x/ansi
	@${TEST} -z ${WRKSRC_charmbracelet_x_ansi} || ${RLN} ${WRKSRC_charmbracelet_x_ansi}/ansi ${WRKSRC}/vendor/github.com/charmbracelet/x
	@${RM} -r ${WRKSRC}/vendor/github.com/charmbracelet/x/cellbuf
	@${TEST} -z ${WRKSRC_charmbracelet_x_cellbuf} || ${RLN} ${WRKSRC_charmbracelet_x_cellbuf}/cellbuf ${WRKSRC}/vendor/github.com/charmbracelet/x
	@${RM} -r ${WRKSRC}/vendor/github.com/charmbracelet/x/term
	@${TEST} -z ${WRKSRC_charmbracelet_x_term} || ${RLN} ${WRKSRC_charmbracelet_x_term}/term ${WRKSRC}/vendor/github.com/charmbracelet/x
	@${MKDIR} ${WRKSRC}/vendor/go.yaml.in/yaml
	@${TEST} ! -e ${WRKSRC}/vendor/gopkg.in/yaml.v3 || ${RLN} ${WRKSRC}/vendor/gopkg.in/yaml.v3 ${WRKSRC}/vendor/go.yaml.in/yaml/v3

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-opentracing:opentracing-go:v1.2.0:opentracing_opentracing_go/vendor/github.com/opentracing/opentracing-go
GH_TUPLE:=	${GH_TUPLE:Nopentracing\:*}

#-moby:moby:v0.7.0:moby_moby_1/vendor/github.com/docker/docker-credential-helpers
#+docker:docker-credential-helpers:v0.7.0:docker_docker_credential_helpers/vendor/github.com/docker/docker-credential-helpers
GH_TUPLE:=	${GH_TUPLE:C|^moby:moby(:.*/docker-credential-helpers)$|docker:docker-credential-helpers\1|}

#-google:go-genproto:3a174f9686a8:google_go_genproto/vendor/google.golang.org/genproto/googleapis/api
#-google:go-genproto:f26f9409b101:google_go_genproto_1/vendor/google.golang.org/genproto/googleapis/rpc
#+google:go-genproto:3a174f9686a8:google_go_genproto/vendor/google.golang.org/genproto/googleapis/api
#+google:go-genproto:f26f9409b101:google_go_genproto_1/vendor/google.golang.org/genproto/googleapis/rpc
GH_TUPLE:=	${GH_TUPLE:C,(_.)?/vendor/google\.golang\.org/genproto/googleapis/(api|rpc),_\2,}

# charmbracelet:x:v0.0.13:charmbracelet_x_cellbuf/vendor/github.com/charmbracelet/x/cellbuf
# charmbracelet:x:v0.2.1:charmbracelet_x_term/vendor/github.com/charmbracelet/x/term
#-charmbracelet:x:v0.8.0:charmbracelet_x_ansi/vendor/github.com/charmbracelet/x/ansi
#+charmbracelet:x:ansi/v0.8.0:charmbracelet_x_ansi
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/charmbracelet/x/(ansi\|cellbuf\|term)|\1\3/\2|}

#-charmbracelet:x:821143405392:charmbracelet_x_slice/vendor/github.com/charmbracelet/x/exp/slice
#-charmbracelet:x:821143405392:charmbracelet_x_strings/vendor/github.com/charmbracelet/x/exp/strings
#+charmbracelet:x:821143405392:charmbracelet_x/vendor/github.com/charmbracelet/x
GH_TUPLE:=	${GH_TUPLE:C|(.*:charmbracelet_x)_.*(/vendor/github.com/charmbracelet/x)/exp/.*|\1\2|:u}

#-containerd:stargz-snapshotter:v0.16.3:containerd_stargz_snapshotter_estargz/vendor/github.com/containerd/stargz-snapshotter/estargz
#+containerd:stargz-snapshotter:v0.16.3:containerd_stargz_snapshotter_estargz/vendor/github.com/containerd/stargz-snapshotter
GH_TUPLE:=	${GH_TUPLE:S|/estargz||1}

#-go-openapi:swag:v0.24.0:go_openapi_swag_cmdutils/vendor/github.com/go-openapi/swag/cmdutils
#-go-openapi:swag:v0.24.0:go_openapi_swag_conv/vendor/github.com/go-openapi/swag/conv
#-go-openapi:swag:v0.24.0:go_openapi_swag_fileutils/vendor/github.com/go-openapi/swag/fileutils
#-go-openapi:swag:v0.24.0:go_openapi_swag_jsonname/vendor/github.com/go-openapi/swag/jsonname
#-go-openapi:swag:v0.24.0:go_openapi_swag_jsonutils/vendor/github.com/go-openapi/swag/jsonutils
#-go-openapi:swag:v0.24.0:go_openapi_swag_loading/vendor/github.com/go-openapi/swag/loading
#-go-openapi:swag:v0.24.0:go_openapi_swag_mangling/vendor/github.com/go-openapi/swag/mangling
#-go-openapi:swag:v0.24.0:go_openapi_swag_netutils/vendor/github.com/go-openapi/swag/netutils
#-go-openapi:swag:v0.24.0:go_openapi_swag_stringutils/vendor/github.com/go-openapi/swag/stringutils
#-go-openapi:swag:v0.24.0:go_openapi_swag_typeutils/vendor/github.com/go-openapi/swag/typeutils
#-go-openapi:swag:v0.24.0:go_openapi_swag_yamlutils/vendor/github.com/go-openapi/swag/yamlutils
# go-openapi:swag:v0.24.1:go_openapi_swag/vendor/github.com/go-openapi/swag
GH_TUPLE:=	${GH_TUPLE:N*/go-openapi/swag/*}

.sinclude "${.CURDIR}/generated.deps.mk"
.include <bsd.port.mk>
