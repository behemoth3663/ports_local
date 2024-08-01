PORTNAME=	tfupdate
PORTVERSION=	0.8.4
DISTVERSIONPREFIX=	v
#ORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Update version constraints in Terraform configurations
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules,1.22

USE_GITHUB=	yes
GH_ACCOUNT=	minamijoyo

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static'"

pre-deps pre-deps-all post-deps-all: .PHONY

.include "/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
