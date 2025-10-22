PORTNAME=	gofumpt
PORTVERSION=	0.9.2
DISTVERSIONPREFIX=	v
CATEGORIES=	local devel

MAINTAINER=	none@freebsd.org
COMMENT=	Stricter gofmt
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	BSD3CLAUSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	mvdan

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X main.version=${PORTVERSION}"

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

post-deps-all: .PHONY
	test -d ${PATCHDIR}_ && ${MV} ${PATCHDIR}_ ${PATCHDIR} || :

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
