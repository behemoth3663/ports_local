PORTNAME=	aws-iam-authenticator
PORTVERSION=	0.7.11
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local security
PKGNAMESUFFIX=	-local

MAINTAINER=	danilo@FreeBSD.org
COMMENT=	Use AWS IAM credentials to authenticate to a Kubernetes cluster
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

BUILD_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	kubernetes-sigs

GO_TARGET=	./cmd/${PORTNAME}

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION} | \
		${SED} -rn 's/(.{7}).*/\1/p'

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X 'sigs.k8s.io/${PORTNAME}/pkg.Version=${PORTVERSION}' \
	-X 'sigs.k8s.io/${PORTNAME}/pkg.BuildDate=${%FT%TZ:L:gmtime}' \
	-X 'sigs.k8s.io/${PORTNAME}/pkg.CommitID=${GIT_COMMIT:sh}'"
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

PLIST_FILES=	bin/${PORTNAME}

M2T_OFFLINE=	true
.export	M2T_OFFLINE

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@${TEST} ! -d ${WRKSRC}/vendor/github.com || ${LN} -s ${WRKSRC}/src/github.com/${GH_ACCOUNT} ${WRKSRC}/vendor/github.com
	@${MKDIR} ${WRKSRC}/vendor/go.yaml.in/yaml
	@${TEST} ! -e ${WRKSRC}/vendor/gopkg.in/yaml.v2 || ${RLN} ${WRKSRC}/vendor/gopkg.in/yaml.v2 ${WRKSRC}/vendor/go.yaml.in/yaml/v2

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=kubernetes GHT_PROJ=klog GHT_PATH=k8s.io/klog/v2 )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=kubernetes-sigs GHT_PROJ=json GHT_PATH=sigs.k8s.io/json )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=kubernetes-sigs GHT_PROJ=structured-merge-diff GHT_PATH=sigs.k8s.io/structured-merge-diff/v6 )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=kubernetes-sigs GHT_PROJ=randfill GHT_PATH=sigs.k8s.io/randfill )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=go-yaml GHT_PROJ=yaml GHT_PATH=go.yaml.in/yaml/v3 )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:space:]]*#[[:space:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:space:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

post-deps-all: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${MV} ${PATCHDIR}_ ${PATCHDIR} || :

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

GH_TUPLE:=	${GH_TUPLE:N*/vendor/github.com/aws/aws-sdk-go-v2/*}
GH_TUPLE:=	${GH_TUPLE:N*/vendor/github.com/go-openapi/swag/*}

#-go-yaml:yaml:v2.4.0:go_yaml_yaml/vendor/gopkg.in/yaml.v2
#+go-yaml:yaml:v2.4.0:go_yaml_yaml_2/vendor/gopkg.in/yaml.v2
GH_TUPLE:=	${GH_TUPLE:C,^(go-yaml:yaml:v2[^:]+:[^/]+),\1_2,1}

.sinclude "${.CURDIR}/generated.deps.mk"

#-go-yaml:yaml:v3.0.4:go_yaml_yaml/vendor/go.yaml.in/yaml/v3
#+go-yaml:yaml:v3.0.1:go_yaml_yaml_3/vendor/go.yaml.in/yaml/v3
GH_TUPLE:=	${GH_TUPLE:C,^(go-yaml:yaml:)v3[^:]+:([^/]+),\1v3.0.1:\2_3,1}

.include <bsd.port.mk>
