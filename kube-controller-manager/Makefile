PORTNAME=	kube-controller-manager
PORTVERSION=	1.25.3
DISTVERSIONPREFIX=	v
CATEGORIES=	local net

MAINTAINER=	loic.blot@unix-experience.fr
COMMENT=	Kubernetes controller manager daemon
WWW=		https://github.com/kubernetes/kubernetes

LICENSE=	APACHE20

RUN_DEPENDS=	git:devel/git

USES=		cpe go:modules
CPE_VENDOR=	jenkins
CPE_PRODUCT=	kubernetes

USE_GITHUB=	yes
GH_ACCOUNT=	kubernetes
GH_PROJECT=	kubernetes

GO_PKGNAME=	k8s.io/kubernetes
GO_TARGET=	${GO_PKGNAME}/cmd/${PORTNAME}

GH_COMMIT=	git ls-remote --tags \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}^{}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
		-X k8s.io/kubernetes/pkg/version.buildDate=${%FT%TZ:L:gmtime} \
		-X k8s.io/kubernetes/pkg/version.gitTreeState= \
		-X k8s.io/kubernetes/pkg/version.gitCommit=${:!${GH_COMMIT}!:[1]} \
		-X k8s.io/kubernetes/pkg/version.gitMajor=${PORTVERSION:R:R} \
		-X k8s.io/kubernetes/pkg/version.gitMinor=${PORTVERSION:R:E} \
		-X k8s.io/kubernetes/pkg/version.gitVersion=${DISTVERSIONPREFIX}${PORTVERSION}"

PLIST_FILES=	bin/${PORTNAME}

KUBERNETES_API_VERSION=	v0.21.5

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

#do not try to avoid `test -z ''` with `.if !empty()` - behaviour is weird
_OPTIONS_extract+=	600:pre-post-extract
pre-post-extract: .PHONY
	@test -z '${WRKSRC_docopt_docopt_go}' || \
		${MV} ${WRKSRC_docopt_docopt_go:S/docopt-go/docopt.go/} ${WRKSRC_docopt_docopt_go}

.include <bsd.port.mk>
