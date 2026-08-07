PORTNAME=	docker
PORTVERSION=	29.7.2
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils
PKGNAMESUFFIX=	-local

MAINTAINER=	ports@FreeBSD.org
COMMENT=	Open-source application container engine
WWW=		https://www.docker.com/

LICENSE=	APACHE20

RUN_DEPENDS=	git:devel/git

USES=		cpe go

PLIST_FILES=	bin/docker

USE_GITHUB=	yes
GH_ACCOUNT=	docker
GH_PROJECT=	cli

GO_TARGET=	./cmd/${PORTNAME}

COMMIT_CMD=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONFULL}^{} | cut -c1-7

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w"
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

post-extract: .PHONY
#	${RM} -r ${WRKSRC}/vendor/go.opentelemetry.io
	${MKDIR} ${WRKSRC}/vendor/github.com/docker/cli
	${RLN} ${WRKSRC}/cli ${WRKSRC}/vendor/github.com/docker/cli
	${RLN} ${WRKSRC}/cli-plugins ${WRKSRC}/vendor/github.com/docker/cli
	${RLN} ${WRKSRC}/cmd ${WRKSRC}/vendor/github.com/docker/cli
	${RLN} ${WRKSRC}/internal ${WRKSRC}/vendor/github.com/docker/cli
	${RLN} ${WRKSRC}/opts ${WRKSRC}/vendor/github.com/docker/cli
	${RLN} ${WRKSRC}/pkg ${WRKSRC}/vendor/github.com/docker/cli
	${RLN} ${WRKSRC}/templates ${WRKSRC}/vendor/github.com/docker/cli
	${REINPLACE_CMD} -r \
		-e 's/^([[:blank:]]*Version[[:blank:]]*=[[:blank:]]*).*/\1"${PORTVERSION}"/' \
		-e "s/^([[:blank:]]*GitCommit[[:blank:]]*=[[:blank:]]*).*/\1\"$$(${COMMIT_CMD})\"/" \
		-e 's/^([[:blank:]]*BuildTime[[:blank:]]*=[[:blank:]]*).*/\1"${%FT%TZ:L:gmtime}"/' \
	${WRKSRC}/cli/version/version.go

.include <bsd.port.mk>
