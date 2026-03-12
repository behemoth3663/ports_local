PORTNAME=	tfvar
PORTVERSION=	0.11.0
DISTVERSIONPREFIX=	v
#ORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Terraform's variable definitions template generator
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	shihanng

PLIST_FILES=	bin/${PORTNAME}

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X main.version=${DISTVERSIONFULL}"
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

post-extract: .PHONY
	${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

pre-deps post-deps-all: .PHONY
pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
