PORTNAME=	golangci-lint
PORTVERSION=	2.12.1
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local devel
PKGNAMESUFFIX=	-local

MAINTAINER=	lcook@FreeBSD.org
COMMENT=	Linters Runner for Go
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	GPLv3

BUILD_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	golangci

USE_GITLAB=	nodefault

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONFULL}

GO_TARGET=	./cmd/${PORTNAME}

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X main.version=${PORTVERSION} \
	-X main.commit=${GIT_COMMIT:sh:C/(.{7}).*/\\1/W} \
	-X main.date=${%FT%TZ:L:gmtime}"
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

PLIST_FILES=	bin/${PORTNAME}

PORTEXAMPLES=	.golangci.reference.yml

OPTIONS_DEFINE=	EXAMPLES

post-install-EXAMPLES-on:
	@${MKDIR} ${STAGEDIR}${EXAMPLESDIR}
	${INSTALL_DATA} ${WRKSRC}/${PORTEXAMPLES} ${STAGEDIR}${EXAMPLESDIR}

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@${RM} -r ${WRKSRC}/vendor/github.com/quasilyte/go-ruleguard/dsl
	@${TEST} -z ${WRKSRC_quasilyte_go_ruleguard_dsl} || ${RLN} ${WRKSRC_quasilyte_go_ruleguard_dsl}/dsl ${WRKSRC}/vendor/github.com/quasilyte/go-ruleguard
	@${RM} -r ${WRKSRC}/vendor/github.com/charmbracelet/x/cellbuf
	@${TEST} -z ${WRKSRC_charmbracelet_x_cellbuf} || ${RLN} ${WRKSRC_charmbracelet_x_cellbuf}/cellbuf ${WRKSRC}/vendor/github.com/charmbracelet/x
	@${RM} -r ${WRKSRC}/vendor/github.com/charmbracelet/x/term
	@${TEST} -z ${WRKSRC_charmbracelet_x_term} || ${RLN} ${WRKSRC_charmbracelet_x_term}/term ${WRKSRC}/vendor/github.com/charmbracelet/x
	@${RM} -r ${WRKSRC}/vendor/github.com/charmbracelet/x/termios
	@${TEST} -z ${WRKSRC_charmbracelet_x_termios} || ${RLN} ${WRKSRC_charmbracelet_x_termios}/termios ${WRKSRC}/vendor/github.com/charmbracelet/x
	@${RM} -r ${WRKSRC}/vendor/github.com/charmbracelet/x/windows
	@${TEST} -z ${WRKSRC_charmbracelet_x_windows} || ${RLN} ${WRKSRC_charmbracelet_x_windows}/windows ${WRKSRC}/vendor/github.com/charmbracelet/x
	@${MKDIR} ${WRKSRC}/vendor/charm.land/lipgloss
	@${TEST} -z ${WRKSRC_charmbracelet_lipgloss_charm} || ${RLN} ${WRKSRC_charmbracelet_lipgloss_charm} ${WRKSRC}/vendor/charm.land/lipgloss/v2
	@${MKDIR} ${WRKSRC}/vendor/codeberg.org/chavacava
	@${TEST} -z ${WRKSRC_chavacava_garif} || ${RLN} ${WRKSRC_chavacava_garif:H}/garif ${WRKSRC}/vendor/codeberg.org/chavacava
	@${MKDIR} ${WRKSRC}/vendor/codeberg.org/polyfloyd
	@${TEST} -z ${WRKSRC_polyfloyd_go_errorlint} || ${RLN} ${WRKSRC_polyfloyd_go_errorlint:H}/go-errorlint ${WRKSRC}/vendor/codeberg.org/polyfloyd
	@${MKDIR} ${WRKSRC}/vendor/go.yaml.in/yaml
	@${TEST} ! -e ${WRKSRC}/vendor/gopkg.in/yaml.v3 || ${RLN} ${WRKSRC}/vendor/gopkg.in/yaml.v3 ${WRKSRC}/vendor/go.yaml.in/yaml/v3

M2T_OFFLINE=	true
.export	M2T_OFFLINE

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=go-simpler GHT_PROJ=musttag GHT_PATH=go-simpler.org/musttag )
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=go-simpler GHT_PROJ=sloglint GHT_PATH=go-simpler.org/sloglint )
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=charmbracelet GHT_PROJ=bubbles GHT_PATH=charm.land/bubbles/v2 )
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=charmbracelet GHT_PROJ=lipgloss GHT_PATH=charm.land/lipgloss/v2 )
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=chavacava GHT_PROJ=garif GHT_PATH=codeberg.org/chavacava/garif )
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=leighmcculloch GHT_PROJ=gocheckcompilerdirectives GHT_PATH=4d63.com/gocheckcompilerdirectives )
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=leighmcculloch GHT_PROJ=gochecknoglobals GHT_PATH=4d63.com/gochecknoglobals )
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=polyfloyd GHT_PROJ=go-errorlint GHT_PATH=codeberg.org/polyfloyd/go-errorlint )
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=Crocmagnon GHT_PROJ=fatcontext GHT_PATH=go.augendre.info/fatcontext )
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=Crocmagnon GHT_PROJ=arangolint GHT_PATH=go.augendre.info/arangolint )
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=GaijinEntertainment GHT_PROJ=golib GHT_PATH=dev.gaijin.team/go/golib )
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=GaijinEntertainment GHT_PROJ=go-exhaustruct GHT_PATH=dev.gaijin.team/go/exhaustruct/v4 )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:blank:]]*#[[:blank:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:blank:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

post-deps-all: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${MV} ${PATCHDIR}_ ${PATCHDIR} || :

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-quasilyte:go-ruleguard:    v0.3.22:quasilyte_go_ruleguard_dsl/vendor/github.com/quasilyte/go-ruleguard/dsl
#+quasilyte:go-ruleguard:dsl/v0.3.22:quasilyte_go_ruleguard_dsl/vendor/github.com/quasilyte/go-ruleguard
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/quasilyte/go-ruleguard/(dsl)|\1\3/\2|1}

#-quasilyte:regex:b3f0c404a727:quasilyte_regex_syntax/vendor/github.com/quasilyte/regex/syntax
#+quasilyte:regex:b3f0c404a727:quasilyte_regex_syntax/vendor/github.com/quasilyte/regex
GH_TUPLE:=	${GH_TUPLE:C|(.*/quasilyte/regex)/syntax$|\1|1}

#-charmbracelet:x:v0.1.4:charmbracelet_x_ansi/vendor/github.com/charmbracelet/x/ansi
#+charmbracelet:x:ansi/v0.1.4:charmbracelet_x_ansi/vendor/github.com/charmbracelet/x
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+):(.*/charmbracelet/x)/(ansi)$|\1:\3/\2|1}

#-charmbracelet:x:v0.2.1:charmbracelet_x_term/vendor/github.com/charmbracelet/x/term
#+charmbracelet:x:term/v0.2.1:charmbracelet_x_term
GH_TUPLE:=	${GH_TUPLE:C,([^:]+:[^:]+:)([^:]+[^/]+).*/charmbracelet/x/(term|termios|windows)$,\1\3/\2,}

.sinclude "${.CURDIR}/generated.deps.mk"

#-chavacava:garif:v0.2.0:chavacava_garif/vendor/codeberg.org/chavacava/garif
#+chavacava:garif:v0.2.0:chavacava_garif
GH_TUPLE:=	${GH_TUPLE:C|/vendor/codeberg\.org/.*||}

.include <bsd.port.mk>
