PORTNAME=	golangci-lint
PORTVERSION=	1.64.7
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local devel

MAINTAINER=	lcook@FreeBSD.org
COMMENT=	Linters Runner for Go
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	GPLv3
LICENSE_FILE=	${WRKSRC}/LICENSE

BUILD_DEPENDS=	git:devel/git

USES=		go:1.24,modules

USE_GITHUB=	yes
GH_ACCOUNT=	golangci

USE_GITLAB=	nodefault

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION} | \
		${SED} -rn 's/(.{7}).*/\1/p'

GO_TARGET=	./cmd/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X main.version=${PORTVERSION} \
	-X main.commit=$$(${GIT_COMMIT}) \
	-X main.date=${%FT%TZ:L:gmtime}"

PLIST_FILES=	bin/${PORTNAME}

PORTEXAMPLES=	.golangci.reference.yml

OPTIONS_DEFINE=	EXAMPLES

post-install-EXAMPLES-on:
	@${MKDIR} ${STAGEDIR}${EXAMPLESDIR}
	${INSTALL_DATA} ${WRKSRC}/${PORTEXAMPLES} \
		${STAGEDIR}${EXAMPLESDIR}

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@test -z '${WRKSRC_quasilyte_go_ruleguard}' || ${RLN} ${WRKSRC_quasilyte_go_ruleguard} ${WRKSRC}/vendor/github.com/quasilyte/go-ruleguard

M2T_OFFLINE=	true
.export	M2T_OFFLINE

pre-deps: .PHONY
	if [ -d '${PATCHDIR}_' ]; then \
		${RM} -r ${PATCHDIR}; \
	else \
		${MV} ${PATCHDIR} ${PATCHDIR}_; \
	fi

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=go-simpler GHT_PROJ=musttag GHT_PATH=go-simpler.org/musttag )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=go-simpler GHT_PROJ=sloglint GHT_PATH=go-simpler.org/sloglint )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=leighmcculloch GHT_PROJ=gocheckcompilerdirectives GHT_PATH=4d63.com/gocheckcompilerdirectives )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=leighmcculloch GHT_PROJ=gochecknoglobals GHT_PATH=4d63.com/gochecknoglobals )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:space:]]*#[[:space:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:space:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

post-deps-all: .PHONY
	test -d '${PATCHDIR}_' && ${MV} ${PATCHDIR}_ ${PATCHDIR} || :

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-quasilyte:go-ruleguard:0fe6f58b47b1:quasilyte_go_ruleguard/vendor/github.com/quasilyte/go-ruleguard
#+quasilyte:go-ruleguard:0fe6f58b47b1:quasilyte_go_ruleguard
GH_TUPLE:=	${GH_TUPLE:S|/vendor/github.com/quasilyte/go-ruleguard$||1}

#-quasilyte:go-ruleguard:    v0.3.22:quasilyte_go_ruleguard_dsl/vendor/github.com/quasilyte/go-ruleguard/dsl
#+quasilyte:go-ruleguard:dsl/v0.3.22:quasilyte_go_ruleguard_dsl/vendor/github.com/quasilyte/go-ruleguard
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+/vendor/github.com/quasilyte/go-ruleguard)/(dsl)|\1\3/\2|}

#-quasilyte:regex:b3f0c404a727:quasilyte_regex_syntax/vendor/github.com/quasilyte/regex/syntax
#+quasilyte:regex:b3f0c404a727:quasilyte_regex_syntax/vendor/github.com/quasilyte/regex
GH_TUPLE:=	${GH_TUPLE:C|(.*/quasilyte/regex)/syntax$|\1|1}

.sinclude "${.CURDIR}/generated.deps.mk"

.include <bsd.port.mk>
