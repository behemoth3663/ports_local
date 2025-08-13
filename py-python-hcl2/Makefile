PORTNAME=	python-hcl2
PORTVERSION=	6.1.0
CATEGORIES=	devel python
MASTER_SITES=	PYPI
PKGNAMEPREFIX=	${PYTHON_PKGNAMEPREFIX}

MAINTAINER=	none@freebsd.org
COMMENT=	Python HCL2 parser (HashiCorp Configuration Language)
WWW=		https://github.com/amplify-education/python-hcl2

LICENSE=	MIT

USES=		python
USE_PYTHON=	autoplist concurrent pep517

BUILD_DEPENDS=	${PYTHON_PKGNAMEPREFIX}setuptools>=0:devel/py-setuptools@${PY_FLAVOR}

RUN_DEPENDS=	${PYTHON_PKGNAMEPREFIX}lark>=0:local/py-lark@${PY_FLAVOR}

#NO_ARCH=	yes
#GH_ACCOUNT=	${WWW:S|/| |g:[-2]}
#GH_PROJECT=	${PORTNAME}

.include <bsd.port.mk>
