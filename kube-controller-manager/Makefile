PORTNAME=	kube-controller-manager
PORTVERSION=	1.25.4
DISTVERSIONPREFIX=	v
CATEGORIES=	local net

MAINTAINER=	loic.blot@unix-experience.fr
COMMENT=	Kubernetes controller manager daemon
WWW=		https://github.com/kubernetes/kubernetes

LICENSE=	APACHE20

BUILD_DEPENDS=	git:devel/git

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
		refs/tags/${GH_PROJECT}-${PORTVERSION}^{}

GO_BUILDFLAGS=	-ldflags="-s -w -extldflags '-static' \
		-X k8s.io/component-base/version.buildDate=${%FT%TZ:L:gmtime} \
		-X k8s.io/component-base/version.gitTreeState= \
		-X k8s.io/component-base/version.gitCommit=${:!${GH_COMMIT}!:[1]} \
		-X k8s.io/component-base/version.gitMajor=${PORTVERSION:R:R} \
		-X k8s.io/component-base/version.gitMinor=${PORTVERSION:R:E} \
		-X k8s.io/component-base/version.gitVersion=${DISTVERSIONPREFIX}${PORTVERSION} \
		-X k8s.io/client-go/pkg/version.gitVersion=${DISTVERSIONPREFIX}${PORTVERSION}"

PLIST_FILES=	bin/${PORTNAME}

#KUBERNETES_API_VERSION!=git ls-remote --tags \
#		https://github.com/kubernetes/api | \
#		sed -rn 's|^.*[[:space:]]refs/tags/(${DISTVERSIONPREFIX}[.0-9]+)\^\{\}[[:space:]]*$$|\1|p' | \
#		sort --version-sort --reverse | head -1
KUBERNETES_API_VERSION=	v0.25.4

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

#do not try to avoid `test -z ''` with `.if !empty()` - behaviour is weird
_OPTIONS_extract+=	600:pre-post-extract
pre-post-extract: .PHONY
	@test -z '${WRKSRC_docopt_docopt_go}' || \
		${MV} ${WRKSRC_docopt_docopt_go:S/docopt-go/docopt.go/} ${WRKSRC_docopt_docopt_go}

.include <bsd.port.mk>
