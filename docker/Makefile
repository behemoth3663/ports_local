# Created by: Nikolai Lifanov <lifanov@FreeBSD.org>

PORTNAME=		docker
PORTVERSION=		19.03.15
DISTVERSIONPREFIX=	v
CATEGORIES=		local sysutils

MAINTAINER=	ports@FreeBSD.org
COMMENT=	Open-source application container engine

LICENSE=	APACHE20

RUN_DEPENDS=	git:devel/git

USES=		cpe go

PLIST_FILES=	bin/docker

USE_GITHUB=	yes
GH_PROJECT=	docker-ce

GO_PKGNAME=	github.com/${PORTNAME}
GO_TARGET=	${GO_PKGNAME}/cli/cmd/docker

GH_COMMIT=	git ls-remote --tags \
		https://github.com/docker/docker-ce \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}

GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X 'github.com/docker/cli/cli/version.Version=${PORTVERSION}' \
	-X 'github.com/docker/cli/cli/version.GitCommit=${:!${GH_COMMIT}!:[1]:C/(.{7}).*/\1/}' \
	-X 'github.com/docker/cli/cli/version.BuildTime=${%FT%TZ:L:gmtime}'"

pre-build: .PHONY
		${FIND} ${GO_WRKSRC}/components -mindepth 1 -maxdepth 1 -type d -exec ${MV} {} ${GO_WRKSRC} \;
		${CP} ${FILESDIR}/default_store_freebsd.go ${GO_WRKSRC}/cli/cli/config/credentials/

.include <bsd.port.mk>
