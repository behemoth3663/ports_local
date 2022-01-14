#

PORTNAME=	ssh-manager
DISTVERSIONPREFIX=	v
DISTVERSION=	0.13.0
CATEGORIES=	local security

MAINTAINER=	none@freebsd.org
COMMENT=	SSH Key Manager for 1Password, Bitwarden and AWS S3

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=	go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	omegion

GO_BUILDFLAGS+=	-ldflags="-s -w \
		-X 'github.com/${GH_ACCOUNT}/${PORTNAME}/internal/info.Version=${DISTVERSIONPREFIX}${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
