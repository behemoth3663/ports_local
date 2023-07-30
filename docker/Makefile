# Created by: Nikolai Lifanov <lifanov@FreeBSD.org>

PORTNAME=		docker
PORTVERSION=		19.03.15
DISTVERSIONPREFIX=	v
CATEGORIES=		local sysutils

MAINTAINER=	ports@FreeBSD.org
COMMENT=	Open-source application container engine
WWW=		https://www.docker.com/

LICENSE=	APACHE20

BUILD_DEPENDS=	git:devel/git

USES=		cpe go

PLIST_FILES=	bin/docker

USE_GITHUB=	yes
GH_PROJECT=	docker-ce

GO_PKGNAME=	github.com/${PORTNAME}
GO_TARGET=	${GO_PKGNAME}/cli/cmd/docker

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION} | \
		${SED} -rn 's/(.{7}).*/\1/p'

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X 'github.com/docker/cli/cli/version.Version=${PORTVERSION}' \
	-X 'github.com/docker/cli/cli/version.GitCommit=$$(${GIT_COMMIT})' \
	-X 'github.com/docker/cli/cli/version.BuildTime=${%FT%TZ:L:gmtime}'"

pre-build: .PHONY
		${FIND} ${GO_WRKSRC}/components -mindepth 1 -maxdepth 1 -type d -exec ${MV} {} ${GO_WRKSRC} \;
		${CP} ${FILESDIR}/default_store_freebsd.go ${GO_WRKSRC}/cli/cli/config/credentials/

.include <bsd.port.mk>
