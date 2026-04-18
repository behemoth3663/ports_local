PORTNAME=	hclgrep
PORTVERSION=	20240307
DISTVERSIONPREFIX=	v
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Syntax based grep for HCL(v2)
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	BSD3CLAUSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	magodo

GH_TAGNAME=	00059cb

PLIST_FILES=	bin/${PORTNAME}

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

post-deps-all: .PHONY
	${TEST} ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
