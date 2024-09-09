PORTNAME=	aws-iam-authenticator
PORTVERSION=	0.6.26
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local security

MAINTAINER=	danilo@FreeBSD.org
COMMENT=	Use AWS IAM credentials to authenticate to a Kubernetes cluster
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

BUILD_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	kubernetes-sigs

GO_TARGET=	./cmd/${PORTNAME}

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION} | \
		${SED} -rn 's/(.{7}).*/\1/p'

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X 'sigs.k8s.io/${PORTNAME}/pkg.Version=${PORTVERSION}' \
	-X 'sigs.k8s.io/${PORTNAME}/pkg.BuildDate=${%FT%TZ:L:gmtime}' \
	-X 'sigs.k8s.io/${PORTNAME}/pkg.CommitID=$$(${GIT_COMMIT})'"

PLIST_FILES=	bin/${PORTNAME}

_OPTIONS_extract=	750:post-post-extract
post-post-extract: .PHONY
	@test -d ${WRKSRC}/vendor/github.com && ${LN} -s ${WRKSRC}/src/github.com/${GH_ACCOUNT} ${WRKSRC}/vendor/github.com/ || :

pre-deps pre-deps-all post-deps-all: .PHONY

.include "/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
