# Created by: Nikolai Lifanov <lifanov@FreeBSD.org>

PORTNAME=		docker
PORTVERSION=		19.03.15
DISTVERSIONPREFIX=	v
CATEGORIES=		sysutils

MAINTAINER=		ports@FreeBSD.org
COMMENT=		Open-source application container engine

LICENSE=		APACHE20

PLIST_FILES=		bin/docker

USE_GITHUB=		yes
GH_PROJECT=		docker-ce

USES=	cpe go

GO_PKGNAME=	github.com/${PORTNAME}
GO_TARGET=	${GO_PKGNAME}/cli/cmd/docker

pre-build: .PHONY
		find ${GO_WRKSRC}/components -mindepth 1 -maxdepth 1 -type d -exec ${MV} {} ${GO_WRKSRC} \;
		${CP} ${FILESDIR}/default_store_freebsd.go ${GO_WRKSRC}/cli/cli/config/credentials/

.include <bsd.port.mk>
