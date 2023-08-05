PORTNAME=	tfupdate
PORTVERSION=	0.7.2
DISTVERSIONPREFIX=	v
#ORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Update version constraints in Terraform configurations
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	minamijoyo

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static'"

PLIST_FILES=	bin/${PORTNAME}

post-post-extract: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
