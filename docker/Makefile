PORTNAME=		docker
PORTVERSION=		29.1.5
DISTVERSIONPREFIX=	v
CATEGORIES=		local sysutils

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

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}^{}

CGO_ENABLED=	0

post-extract: .PHONY
	# ${RM} -r ${WRKSRC}/vendor/go.opentelemetry.io
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
		-e 's/^([[:blank:]]*GitCommit[[:blank:]]*=[[:blank:]]*).*/\1"${GIT_COMMIT:sh:C/(.{7}).*/\\1/W}"/' \
		-e 's/^([[:blank:]]*BuildTime[[:blank:]]*=[[:blank:]]*).*/\1"${%FT%TZ:L:gmtime}"/' \
	${WRKSRC}/cli/version/version.go

.include <bsd.port.mk>
