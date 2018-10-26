#!/bin/bash
swagger generate client -f ../third_party/equinix/ecx/ecxfabric-v3-catalog-seller.json go-ecxfabric/seller
# apply o_auth tokentimeout patch
patch models/o_auth_response.go < seller_patch.diff
