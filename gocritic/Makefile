PORTNAME=	gocritic
PORTVERSION=	0.14.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	7
CATEGORIES=	local devel

MAINTAINER=	fuz@FreeBSD.org
COMMENT=	Opinionated Go source code linter for code audit
WWW=		https://go-critic.com/

LICENSE=	MIT

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	go-critic
GH_PROJECT=	go-critic

GO_TARGET=	./cmd/gocritic \
		./cmd/gocritic-analysis

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X main.Version=v${PORTVERSION}"

PLIST_FILES=	bin/gocritic \
		bin/gocritic-analysis

M2T_OFFLINE=	true
.export	M2T_OFFLINE

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@${RM} -r ${WRKSRC}/vendor/github.com/quasilyte/go-ruleguard/dsl
	@test -z '${WRKSRC_quasilyte_go_ruleguard_dsl}' || ${RLN} ${WRKSRC_quasilyte_go_ruleguard_dsl}/dsl ${WRKSRC}/vendor/github.com/quasilyte/go-ruleguard/dsl

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=golang GHT_PROJ=exp GHT_PATH=golang.org/x/exp/typeparams )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:space:]]*#[[:space:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:space:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

post-deps-all: .PHONY
	test -d '${PATCHDIR}_' && ${MV} ${PATCHDIR}_ ${PATCHDIR} || :

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-quasilyte:go-ruleguard:v0.3.22:quasilyte_go_ruleguard_dsl/vendor/github.com/quasilyte/go-ruleguard/dsl
#+quasilyte:go-ruleguard:dsl/v0.3.22:quasilyte_go_ruleguard_dsl
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+)/vendor/github.com/quasilyte/go-ruleguard/(dsl)|\1\3/\2|1}

#-quasilyte:regex:b3f0c404a727:quasilyte_regex_syntax/vendor/github.com/quasilyte/regex/syntax
#+quasilyte:regex:b3f0c404a727:quasilyte_regex_syntax/vendor/github.com/quasilyte/regex
GH_TUPLE:=	${GH_TUPLE:C|(.*/quasilyte/regex)/syntax$|\1|1}

.sinclude "${.CURDIR}/generated.deps.mk"

GH_TUPLE:=	${GH_TUPLE:S|/typeparams||1}

.include <bsd.port.mk>
