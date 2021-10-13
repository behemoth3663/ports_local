#

PORTNAME=	terrascan
PORTVERSION=	1.11.0
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Compliance and security violations in terraform code

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules
GO_TARGET=	./cmd/${PORTNAME}

USE_GITHUB=	yes
USE_GITLAB=	nodefault

GH_ACCOUNT=	accurics

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

_OPTIONS_extract=	701:post-post-extract

post-post-extract: .PHONY
	find ${WRKSRC}/vendor/modernc.org/libc -name '*_linux*.go' -type f | while read src; do \
		dst=$$(echo "$${src}" | sed 's/_linux/_freebsd/'); \
		test -f "$${dst}" || ln -s "$${src}" "$${dst}"; \
	done
	${MV} ${WRKSRC}/vendor/sigs.k8s.io/kustomize/api/api/* ${WRKSRC}/vendor/sigs.k8s.io/kustomize/api || :
	${MV} ${WRKSRC}/vendor/sigs.k8s.io/kustomize/kyaml/kyaml/* ${WRKSRC}/vendor/sigs.k8s.io/kustomize/kyaml || :

.-include "${.CURDIR}/Makefile.deps"

#-		kubernetes-sigs:kustomize:v0.10.15:kubernetes_sigs_kustomize_1/vendor/sigs.k8s.io/kustomize/kyaml \
#-		kubernetes-sigs:kustomize:v0.8.5:kubernetes_sigs_kustomize/vendor/sigs.k8s.io/kustomize/api \
#+		kubernetes-sigs:kustomize:kyaml/v0.10.15:kubernetes_sigs_kustomize_1/vendor/sigs.k8s.io/kustomize/kyaml \
#+		kubernetes-sigs:kustomize:api/v0.8.5:kubernetes_sigs_kustomize/vendor/sigs.k8s.io/kustomize/api \

#pre-build: .PHONY
#	cd ${WRKSRC} && ${GO_CMD} mod tidy -v && ${GO_CMD} mod download

.include <bsd.port.mk>
#.include <bsd.port.pre.mk>
#GO_CMD=	${LOCALBASE}/go-1.16/bin/go
#GO_BUILDFLAGS=	-tags=ooni,nomk,DISABLE_QUIC -ldflags '-s -w -extldflags "-static"'
#GO_BUILDFLAGS:=	${GO_BUILDFLAGS:N-mod=vendor} -mod=mod
#.include <bsd.port.post.mk>
