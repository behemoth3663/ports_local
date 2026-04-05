PORTNAME=	dumbproxy
PORTVERSION=	1.45.2
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local net

MAINTAINER=	none@freebsd.org
COMMENT=	Simple, scriptable, secure HTTP/SOCKS5 forward proxy
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	Snawoot

PLIST_FILES=	bin/${PORTNAME}

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X 'main.version=${PORTVERSION}'"
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

#M2T_OFFLINE=	true
#.export	M2T_OFFLINE

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@${MKDIR} ${WRKSRC}/vendor/codeberg.org/yarmak
	@${TEST} -z ${WRKSRC_yarmak_secache} || ${RLN} ${WRKSRC_yarmak_secache:H}/secache ${WRKSRC}/vendor/codeberg.org/yarmak

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
#	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=yarmak GHT_PROJ=secache GHT_PATH=codeberg.org/yarmak/secache )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:blank:]]*#[[:blank:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:blank:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

post-deps-all: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${MV} ${PATCHDIR}_ ${PATCHDIR} || :

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.sinclude "${.CURDIR}/generated.deps.mk"

#-yarmak:secache:v0.2.4:yarmak_secache/vendor/codeberg.org/yarmak/secache
#+yarmak:secache:v0.2.4:yarmak_secache
GH_TUPLE:=	${GH_TUPLE:C|/vendor/codeberg\.org/.*||}

.include <bsd.port.mk>
