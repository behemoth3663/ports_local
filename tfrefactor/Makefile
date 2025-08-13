PORTNAME=	tfrefactor
PORTVERSION=	0.0.1
DISTVERSIONPREFIX=	release/v
#ORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Update version constraints in Terraform configurations
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	craftvscruft

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static'"

PLIST_FILES=	bin/${PORTNAME}

post-post-extract: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
