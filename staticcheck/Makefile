PORTNAME=	staticcheck
PORTVERSION=	0.8.2
DISTVERSIONPREFIX=	v
CATEGORIES=	local devel

MAINTAINER=	none@freebsd.org
COMMENT=	Advanced Go linter
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	dominikh
GH_PROJECT=	go-tools

GO_TARGET=	./cmd/${PORTNAME}

CGO_ENABLED=	0
GO_LDFLAGS+=	-w
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

PLIST_FILES=	bin/${PORTNAME}

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

post-deps-all: .PHONY
	${TEST} ! -d '${PATCHDIR}_' || ${MV} ${PATCHDIR}_ ${PATCHDIR}

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
