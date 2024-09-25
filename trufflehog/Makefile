PORTNAME=	trufflehog
PORTVERSION=	3.82.4
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	yuri@FreeBSD.org
COMMENT=	Find and verify credentials
WWW=		https://trufflesecurity.com/

LICENSE=	AGPLv3
LICENSE_FILE=	${WRKSRC}/LICENSE

BROKEN_armv7=	compilation fails see details here: https://github.com/golang/go/issues/65187
BROKEN_i386=	compilation fails see details here: https://github.com/golang/go/issues/65187

USES=		go:modules,1.23

GO_MODULE=	github.com/trufflesecurity/trufflehog/v3

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X '${GO_MODULE}/pkg/version.BuildVersion=${DISTVERSIONPREFIX}${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

.include <bsd.port.mk>
