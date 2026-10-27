PORTNAME=	gh
PORTVERSION=	2.103.3
DISTVERSIONPREFIX=	v
PORTREVISION=	1
CATEGORIES=	local devel
PKGNAMESUFFIX=	-local

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
GO_LDFLAGS+=	-w
GO_LDFLAGS+=	-X github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/internal/build.Version=${PORTVERSION}
GO_LDFLAGS+=	-X github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/internal/build.Date=${SOURCE_DATE_EPOCH_CMD:sh}
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

_DOCS=		README.md docs/[!README]*.md
PORTDOCS=	${_DOCS:T}

OPTIONS_DEFINE=		DOCS MANPAGES
OPTIONS_DEFAULT=	MANPAGES
OPTIONS_SUB=		yes

M2T_OFFLINE=	true
.export	M2T_OFFLINE

clean-unused: .PHONY clean patch
	( cd ${WRKSRC} && ${GO_CMD} mod tidy -v )
	${RM} -r ${WRKSRC}/vendor
	( ${MAKE} -C ${.CURDIR} gomod-vendor | \
		${GREP} --extended-regexp --invert-match '^(/|=+>|\*+[[:blank:]]+Error)' | \
		${SED} -r -e '/\r$$/d' -e '/^[[:blank:]]+"message":/d' -e '/\\}/d' -e '/^[[:blank:]]*$$/d' \
		> ${.CURDIR}/Makefile.deps )
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}
	( ${MAKE} -C ${.CURDIR} makesum )

do-build-MANPAGES-on: .PHONY
	(cd ${WRKSRC} && ${SETENV} ${MAKE_ENV} ${GO_ENV} ${GO_CMD} run ${GO_BUILDFLAGS} \
		./cmd/gen-docs --man-page --doc-path docs/man 2>/dev/null)

post-install-DOCS-on: .PHONY
	@${MKDIR} ${STAGEDIR}${DOCSDIR}
	( cd ${WRKSRC} && ${INSTALL_DATA} ${_DOCS} ${STAGEDIR}${DOCSDIR} )

post-install-MANPAGES-on: .PHONY
	${INSTALL_MAN} ${WRKSRC}/docs/man/*.1 ${STAGEDIR}${PREFIX}/share/man/man1

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=imdario GHT_PROJ=mergo GHT_PATH=dario.cat/mergo )
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=charmbracelet GHT_PROJ=bubbles GHT_PATH=charm.land/bubbles/v2 )
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=charmbracelet GHT_PROJ=bubbletea GHT_PATH=charm.land/bubbletea/v2 )
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=charmbracelet GHT_PROJ=huh GHT_PATH=charm.land/huh/v2 )
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=charmbracelet GHT_PROJ=glamour GHT_PATH=charm.land/glamour/v2 )
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=charmbracelet GHT_PROJ=lipgloss GHT_PATH=charm.land/lipgloss/v2 )
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=kubernetes GHT_PROJ=klog GHT_PATH=k8s.io/klog/v2 )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:blank:]]*#[[:blank:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:blank:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

post-deps-all: .PHONY
	${TEST} ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor || :
	@${RM} -r ${WRKSRC}/internal/barista/observability
	@${RM} -r ${WRKSRC}/internal/telemetry
	@${RM} -r ${WRKSRC}/pkg/cmd/send-telemetry
	@${RM} ${WRKSRC}/internal/attachments/telemetry.go
	@${RM} ${WRKSRC}/pkg/cmdutil/telemetry.go
	@${RM} ${WRKSRC}/pkg/cmd/auth/shared/telemetry.go
	@${RM} ${WRKSRC}/vendor/github.com/go-openapi/runtime/client/opentelemetry.go
	@${FIND} ${WRKSRC} -name '*_test.go' -delete
	@${MKDIR} ${WRKSRC}/vendor/google.golang.org/genproto/googleapis
	@${TEST} -z ${WRKSRC_google_go_genproto_api} || ${RLN} ${WRKSRC_google_go_genproto_api}/googleapis/api ${WRKSRC}/vendor/google.golang.org/genproto/googleapis
	@${TEST} -z ${WRKSRC_google_go_genproto_rpc} || ${RLN} ${WRKSRC_google_go_genproto_rpc}/googleapis/rpc ${WRKSRC}/vendor/google.golang.org/genproto/googleapis
	@${MKDIR} ${WRKSRC}/vendor/github.com/charmbracelet/x/exp
	@${TEST} -z ${WRKSRC_charmbracelet_x_ansi} || ${RLN} ${WRKSRC_charmbracelet_x_ansi}/ansi ${WRKSRC}/vendor/github.com/charmbracelet/x
	@${TEST} -z ${WRKSRC_charmbracelet_x_cellbuf} || ${RLN} ${WRKSRC_charmbracelet_x_cellbuf}/cellbuf ${WRKSRC}/vendor/github.com/charmbracelet/x
	@${TEST} -z ${WRKSRC_charmbracelet_x_term} || ${RLN} ${WRKSRC_charmbracelet_x_term}/term ${WRKSRC}/vendor/github.com/charmbracelet/x
	@${TEST} -z ${WRKSRC_charmbracelet_x_termios} || ${RLN} ${WRKSRC_charmbracelet_x_termios}/termios ${WRKSRC}/vendor/github.com/charmbracelet/x
	@${TEST} -z ${WRKSRC_charmbracelet_x_windows} || ${RLN} ${WRKSRC_charmbracelet_x_windows}/windows ${WRKSRC}/vendor/github.com/charmbracelet/x
	@${TEST} -z ${WRKSRC_charmbracelet_x_strings} || ${RLN} ${WRKSRC_charmbracelet_x_strings}/exp/strings ${WRKSRC}/vendor/github.com/charmbracelet/x/exp
	@${TEST} -z ${WRKSRC_charmbracelet_x_strings} || ${RLN} ${WRKSRC_charmbracelet_x_strings}/exp/slice ${WRKSRC}/vendor/github.com/charmbracelet/x/exp
	@${TEST} -z ${WRKSRC_charmbracelet_x_ordered} || ${RLN} ${WRKSRC_charmbracelet_x_ordered}/exp/ordered ${WRKSRC}/vendor/github.com/charmbracelet/x/exp
	@${MKDIR} ${WRKSRC}/vendor/charm.land/lipgloss
	@${TEST} -z ${WRKSRC_charmbracelet_lipgloss_charm} || ${RLN} ${WRKSRC_charmbracelet_lipgloss_charm} ${WRKSRC}/vendor/charm.land/lipgloss/v2
	@${MKDIR} ${WRKSRC}/vendor/go.yaml.in/yaml
	@${TEST} ! -e ${WRKSRC}/vendor/gopkg.in/yaml.v3 || ${RLN} ${WRKSRC}/vendor/gopkg.in/yaml.v3 ${WRKSRC}/vendor/go.yaml.in/yaml/v3

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-opentracing:opentracing-go:v1.2.0:opentracing_opentracing_go/vendor/github.com/opentracing/opentracing-go
GH_TUPLE:=	${GH_TUPLE:Nopentracing\:*}

#-google:go-genproto:3a174f9686a8:google_go_genproto/vendor/google.golang.org/genproto/googleapis/api
#-google:go-genproto:f26f9409b101:google_go_genproto_1/vendor/google.golang.org/genproto/googleapis/rpc
#+google:go-genproto:3a174f9686a8:google_go_genproto_api
#+google:go-genproto:f26f9409b101:google_go_genproto_rpc
GH_TUPLE:=	${GH_TUPLE:C,(_.)?/vendor/google\.golang\.org/genproto/googleapis/(api|rpc),,}

#-charmbracelet:x:v0.8.0:charmbracelet_x_ansi/vendor/github.com/charmbracelet/x/ansi
#+charmbracelet:x:ansi/v0.8.0:charmbracelet_x_ansi
GH_TUPLE:=	${GH_TUPLE:C,:([^:]+:[^:/]+)/.*/charmbracelet/x/(ansi|cellbuf|term|termios|windows|exp/ordered),:\2/\1,}

#-charmbracelet:x:821143405392:charmbracelet_x_slice/vendor/github.com/charmbracelet/x/exp/slice
#-charmbracelet:x:821143405392:charmbracelet_x_strings/vendor/github.com/charmbracelet/x/exp/strings
#+charmbracelet:x:821143405392:charmbracelet_x_strings
GH_TUPLE:=	${GH_TUPLE:C,([^:]+:[^:]+:)([^:]+[^/]+).*/charmbracelet/x/exp/strings,\1\2,1}
GH_TUPLE:=	${GH_TUPLE:N*/charmbracelet/x/exp/slice}

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

#-charmbracelet:lipgloss:v2.0.2:charmbracelet_lipgloss/vendor/charm.land/lipgloss/v2
#+charmbracelet:lipgloss:v2.0.2:charmbracelet_lipgloss_charm
GH_TUPLE:=	${GH_TUPLE:C,([^:]+:[^:]+:)([^:]+[^/]+).*/charm.land/lipgloss/v2,\1\2_charm,1}

.include <bsd.port.pre.mk>

SOURCE_DATE_EPOCH_CMD=	date -ur \
			$$(${GREP} --max-count=1 TIMESTAMP ${DISTINFO_FILE} | ${CUT} -d' ' -f3) \
			'+%FT%TZ'

.include <bsd.port.post.mk>
