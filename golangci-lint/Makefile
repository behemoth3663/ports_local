PORTNAME=	golangci-lint
PORTVERSION=	2.8.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local devel

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
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION} | \
		${SED} -rn 's/(.{7}).*/\1/p'

GO_TARGET=	./cmd/${PORTNAME}

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X main.version=${PORTVERSION} \
	-X main.commit=$$(${GIT_COMMIT}) \
	-X main.date=${%FT%TZ:L:gmtime}"

PLIST_FILES=	bin/${PORTNAME}

PORTEXAMPLES=	.golangci.reference.yml

OPTIONS_DEFINE=	EXAMPLES

post-install-EXAMPLES-on:
	@${MKDIR} ${STAGEDIR}${EXAMPLESDIR}
	${INSTALL_DATA} ${WRKSRC}/${PORTEXAMPLES} ${STAGEDIR}${EXAMPLESDIR}

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@${RM} -r ${WRKSRC}/vendor/github.com/quasilyte/go-ruleguard/dsl
	@test -z ${WRKSRC_quasilyte_go_ruleguard_dsl} || ${RLN} ${WRKSRC_quasilyte_go_ruleguard_dsl}/dsl ${WRKSRC}/vendor/github.com/quasilyte/go-ruleguard
	@${RM} -r ${WRKSRC}/vendor/github.com/charmbracelet/x/cellbuf
	@test -z ${WRKSRC_charmbracelet_x_cellbuf} || ${RLN} ${WRKSRC_charmbracelet_x_cellbuf}/cellbuf ${WRKSRC}/vendor/github.com/charmbracelet/x
	@${RM} -r ${WRKSRC}/vendor/github.com/charmbracelet/x/term
	@test -z ${WRKSRC_charmbracelet_x_term} || ${RLN} ${WRKSRC_charmbracelet_x_term}/term ${WRKSRC}/vendor/github.com/charmbracelet/x
	@${MKDIR} ${WRKSRC}/vendor/codeberg.org/chavacava
	@test -z ${WRKSRC_chavacava_garif} || ${RLN} ${WRKSRC_chavacava_garif:H}/garif ${WRKSRC}/vendor/codeberg.org/chavacava
	@${MKDIR} ${WRKSRC}/vendor/codeberg.org/polyfloyd
	@test -z ${WRKSRC_polyfloyd_go_errorlint} || ${RLN} ${WRKSRC_polyfloyd_go_errorlint:H}/go-errorlint ${WRKSRC}/vendor/codeberg.org/polyfloyd
	@${MKDIR} ${WRKSRC}/vendor/go.yaml.in/yaml
	@${TEST} ! -e ${WRKSRC}/vendor/gopkg.in/yaml.v3 || ${RLN} ${WRKSRC}/vendor/gopkg.in/yaml.v3 ${WRKSRC}/vendor/go.yaml.in/yaml/v3

M2T_OFFLINE=	true
.export	M2T_OFFLINE

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=go-simpler GHT_PROJ=musttag GHT_PATH=go-simpler.org/musttag )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=go-simpler GHT_PROJ=sloglint GHT_PATH=go-simpler.org/sloglint )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=leighmcculloch GHT_PROJ=gocheckcompilerdirectives GHT_PATH=4d63.com/gocheckcompilerdirectives )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=leighmcculloch GHT_PROJ=gochecknoglobals GHT_PATH=4d63.com/gochecknoglobals )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=Crocmagnon GHT_PROJ=fatcontext GHT_PATH=go.augendre.info/fatcontext )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=Crocmagnon GHT_PROJ=arangolint GHT_PATH=go.augendre.info/arangolint )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=chavacava GHT_PROJ=garif GHT_PATH=codeberg.org/chavacava/garif )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=polyfloyd GHT_PROJ=go-errorlint GHT_PATH=codeberg.org/polyfloyd/go-errorlint )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=GaijinEntertainment GHT_PROJ=golib GHT_PATH=dev.gaijin.team/go/golib )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=GaijinEntertainment GHT_PROJ=go-exhaustruct GHT_PATH=dev.gaijin.team/go/exhaustruct/v4 )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:space:]]*#[[:space:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:space:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

post-deps-all: .PHONY
	test -d ${PATCHDIR}_ && ${MV} ${PATCHDIR}_ ${PATCHDIR} || :

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-quasilyte:go-ruleguard:0fe6f58b47b1:quasilyte_go_ruleguard/vendor/github.com/quasilyte/go-ruleguard
#+quasilyte:go-ruleguard:0fe6f58b47b1:quasilyte_go_ruleguard
#GH_TUPLE:=	${GH_TUPLE:S|/vendor/github.com/quasilyte/go-ruleguard$||1}

#-quasilyte:go-ruleguard:    v0.3.22:quasilyte_go_ruleguard_dsl/vendor/github.com/quasilyte/go-ruleguard/dsl
#+quasilyte:go-ruleguard:dsl/v0.3.22:quasilyte_go_ruleguard_dsl/vendor/github.com/quasilyte/go-ruleguard
#GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+/vendor/github.com/quasilyte/go-ruleguard)/(dsl)|\1\3/\2|}
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/quasilyte/go-ruleguard/(dsl)|\1\3/\2|1}

#-		charmbracelet:x:     v0.1.4:charmbracelet_x_ansi/vendor/github.com/charmbracelet/x/ansi
#+		charmbracelet:x:ansi/v0.1.4:charmbracelet_x_ansi/vendor/github.com/charmbracelet/x
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+):(.*/charmbracelet/x)/(ansi)|\1:\3/\2|1}

#-charmbracelet:x:     v0.2.1:charmbracelet_x_term/vendor/github.com/charmbracelet/x/term
#+charmbracelet:x:term/v0.2.1:charmbracelet_x_term
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/charmbracelet/x/(term)|\1\3/\2|1}
GH_TUPLE:=	${GH_TUPLE:S|/vendor/github.com/charmbracelet/x/cellbuf$||1}

#-quasilyte:regex:b3f0c404a727:quasilyte_regex_syntax/vendor/github.com/quasilyte/regex/syntax
#+quasilyte:regex:b3f0c404a727:quasilyte_regex_syntax/vendor/github.com/quasilyte/regex
GH_TUPLE:=	${GH_TUPLE:C|(.*/quasilyte/regex)/syntax$|\1|1}

GH_TUPLE+=	tetafro:godot:v1.5.4:tetafro_godot/vendor/github.com/tetafro/godot

.sinclude "${.CURDIR}/generated.deps.mk"

#-chavacava:garif:v0.2.0:chavacava_garif/vendor/codeberg.org/chavacava/garif
#+chavacava:garif:v0.2.0:chavacava_garif
GH_TUPLE:=	${GH_TUPLE:C|/vendor/codeberg\.org/.*||}

.include <bsd.port.mk>
