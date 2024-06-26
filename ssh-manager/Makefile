PORTNAME=	ssh-manager
DISTVERSIONPREFIX=	v
DISTVERSION=	1.12.3
CATEGORIES=	local security

MAINTAINER=	none@freebsd.org
COMMENT=	SSH Key Manager for 1Password, Bitwarden and AWS S3
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	omegion

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -buildid=${PORTVERSION} -extldflags '-static'"

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
