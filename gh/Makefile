PORTNAME=	gh
PORTVERSION=	2.76.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local devel

MAINTAINER=	lcook@FreeBSD.org
COMMENT=	GitHub's official command line tool
WWW=		https://cli.github.com

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

RUN_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	cli
GH_PROJECT=	${GH_ACCOUNT}

GO_TARGET=	./cmd/${PORTNAME}
GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/internal/build.Version=${PORTVERSION} \
	-X github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/internal/build.Date=${%FT%TZ:L:gmtime}"

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
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=imdario GHT_PROJ=mergo GHT_PATH=dario.cat/mergo )
#	( make -C ${.CURDIR} make-from-deps GHT_ACNT=kubernetes GHT_PROJ=klog GHT_PATH=k8s.io/klog/v2 )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:space:]]*#[[:space:]]*::(v[0-9.]+):.*/(vendor/${GHT_PATH})[[:space:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/}_${GHT_PROJ:tl:S/-/_/}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

post-deps-all: .PHONY
	test ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@${MKDIR} ${WRKSRC}/vendor/google.golang.org/genproto/googleapis
	@test -z '${WRKSRC_google_go_genproto_api}' || ${RLN} ${WRKSRC_google_go_genproto_api}/googleapis/api ${WRKSRC}/vendor/google.golang.org/genproto/googleapis/api
	@test -z '${WRKSRC_google_go_genproto_rpc}' || ${RLN} ${WRKSRC_google_go_genproto_rpc}/googleapis/rpc ${WRKSRC}/vendor/google.golang.org/genproto/googleapis/rpc
	@${RM} ${WRKSRC}/vendor/github.com/go-openapi/runtime/client/opentelemetry.go
	@${RM} ${WRKSRC}/vendor/github.com/go-openapi/runtime/client/opentelemetry_test.go
	@${RM} ${WRKSRC}/vendor/github.com/go-openapi/runtime/client/opentracing.go
	@${RM} ${WRKSRC}/vendor/github.com/go-openapi/runtime/client/opentracing_test.go
	@${RM} -r ${WRKSRC}/vendor/github.com/charmbracelet/x/ansi
	@test -z '${WRKSRC_charmbracelet_x_ansi}' || ${RLN} ${WRKSRC_charmbracelet_x_ansi}/ansi ${WRKSRC}/vendor/github.com/charmbracelet/x
	@${RM} -r ${WRKSRC}/vendor/github.com/charmbracelet/x/cellbuf
	@test -z '${WRKSRC_charmbracelet_x_cellbuf}' || ${RLN} ${WRKSRC_charmbracelet_x_cellbuf}/cellbuf ${WRKSRC}/vendor/github.com/charmbracelet/x
	@${RM} -r ${WRKSRC}/vendor/github.com/charmbracelet/x/term
	@test -z '${WRKSRC_charmbracelet_x_term}' || ${RLN} ${WRKSRC_charmbracelet_x_term}/term ${WRKSRC}/vendor/github.com/charmbracelet/x

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-moby:moby:v0.7.0:moby_moby_1/vendor/github.com/docker/docker-credential-helpers
#+docker:docker-credential-helpers:v0.7.0:docker_docker_credential_helpers/vendor/github.com/docker/docker-credential-helpers
GH_TUPLE:=	${GH_TUPLE:C|^moby:moby(:.*/docker-credential-helpers)$|docker:docker-credential-helpers\1|}

#-google:go-genproto:c811ad7063a7:google_go_genproto/vendor/google.golang.org/genproto/googleapis/api
#+google:go-genproto:c811ad7063a7:google_go_genproto_api
#-google:go-genproto:94a12d6c2237:google_go_genproto_1/vendor/google.golang.org/genproto/googleapis/rpc
#+google:go-genproto:94a12d6c2237:google_go_genproto_rpc
GH_TUPLE:=	${GH_TUPLE:C|(:google_go_genproto).*/|\1_|}

# charmbracelet:x:v0.0.13:charmbracelet_x_cellbuf/vendor/github.com/charmbracelet/x/cellbuf
# charmbracelet:x:v0.2.1:charmbracelet_x_term/vendor/github.com/charmbracelet/x/term
#-charmbracelet:x:v0.8.0:charmbracelet_x_ansi/vendor/github.com/charmbracelet/x/ansi
#+charmbracelet:x:ansi/v0.8.0:charmbracelet_x_ansi
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/charmbracelet/x/(ansi\|cellbuf\|term)|\1\3/\2|}

#-charmbracelet:x:212f7b056ed0:charmbracelet_x_strings/vendor/github.com/charmbracelet/x/exp/strings
#+charmbracelet:x:212f7b056ed0:charmbracelet_x_strings/vendor/github.com/charmbracelet/x
GH_TUPLE:=	${GH_TUPLE:S|/charmbracelet/x/exp/strings|/charmbracelet/x|1}

#-		containerd:stargz-snapshotter:v0.16.3:containerd_stargz_snapshotter_estargz/vendor/github.com/containerd/stargz-snapshotter/estargz
#+		containerd:stargz-snapshotter:v0.16.3:containerd_stargz_snapshotter_estargz/vendor/github.com/containerd/stargz-snapshotter
GH_TUPLE:=	${GH_TUPLE:S|/estargz||1}

#-		opentracing:opentracing-go:v1.2.0:opentracing_opentracing_go/vendor/github.com/opentracing/opentracing-go
GH_TUPLE:=	${GH_TUPLE:Nopentracing\:opentracing-go\:*}

.sinclude "${.CURDIR}/generated.deps.mk"
.include <bsd.port.mk>
