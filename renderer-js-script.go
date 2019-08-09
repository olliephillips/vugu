package vugu

// GENERATED FILE, DO NOT EDIT!  See renderer-js-script-maker.go

const jsHelperScript = "\n(function() {\n\n\tif (window.vuguRender) { return; } // only once\n\n    const opcodeEnd = 0         // no more instructions in this buffer\n    // const opcodeClearRefmap = 1 // clear the reference map, all following instructions must not reference prior IDs\n    const opcodeClearEl = 1 // clear the currently selected element\n    // const opcodeSetHTMLRef = 2  // assign ref for html tag\n    // const opcodeSetHeadRef = 3  // assign ref for head tag\n    // const opcodeSetBodyRef = 4  // assign ref for body tag\n    // const opcodeSelectRef = 5   // select element by ref\n\tconst opcodeRemoveOtherAttrs = 5 // remove any elements for the current element that we didn't just set\n    const opcodeSetAttrStr = 6  // assign attribute string to the current selected element\n    const opcodeSelectMountPoint = 7 // selects the mount point element and pushes to the stack - the first time by selector but every subsequent time it will reuse the element from before (because the selector may not match after it's been synced over, it's id etc), also make sure it's of this element name and recreate if so\n\t// const opcodePicardFirstChildElement = 8  // ensure an element first child and push onto element stack\n\t// const opcodePicardFirstChildText    = 9  // ensure a text first child and push onto element stack\n\t// const opcodePicardFirstChildComment = 10 // ensure a comment first child and push onto element stack\n\tconst opcodeSelectParent                   = 11 // pop from the element stack\n\tconst opcodePicardFirstChild = 12  // ensure an element first child and push onto element stack\n\n    const opcodeMoveToFirstChild     = 20 // move node selection to first child (doesn't have to exist)\n\tconst opcodeSetElement           = 21 // assign current selected node as an element of the specified type\n\t// const opcodeSetElementAttr       = 22 // set attribute on current element\n\tconst opcodeSetText              = 23 // assign current selected node as text with specified content\n\tconst opcodeSetComment           = 24 // assign current selected node as comment with specified content\n\tconst opcodeMoveToParent         = 25 // move node selection to parent\n\tconst opcodeMoveToNextSibling    = 26 // move node selection to next sibling (doesn't have to exist)\n\tconst opcodeClearEventListeners  = 27 // remove all event listeners from currently selected element\n\tconst opcodeSetEventListener     = 28 // assign event listener to currently selected element\n\n    // Decoder provides our binary decoding.\n    // Using a class because that's what all the cool JS kids are doing these days.\n    class Decoder {\n\n        constructor(dataView, offset) {\n            this.dataView = dataView;\n            this.offset = offset || 0;\n            return this;\n        }\n\n        // readUint8 reads a single byte, 0-255\n        readUint8() {\n            var ret = this.dataView.getUint8(this.offset);\n            this.offset++;\n            return ret;\n        }\n\n        // readRefToString reads a 64-bit unsigned int ref but returns it as a hex string\n        readRefToString() {\n            // read in two 32-bit parts, BigInt is not yet well supported\n            var ret = this.dataView.getUint32(this.offset).toString(16).padStart(8, \"0\") +\n                this.dataView.getUint32(this.offset + 4).toString(16).padStart(8, \"0\");\n            this.offset += 8;\n            return ret;\n        }\n\n        // readString is 4 bytes length followed by utf chars\n        readString() {\n            var len = this.dataView.getUint32(this.offset);\n            var ret = utf8decoder.decode(new DataView(this.dataView.buffer, this.dataView.byteOffset + this.offset + 4, len));\n            this.offset += len + 4;\n            return ret;\n        }\n\n    }\n\n    let utf8decoder = new TextDecoder();\n\n\twindow.vuguRender = function(buffer) { \n        \n        // NOTE: vuguRender must not automatically reset anything between calls.\n        // Since a series of instructions might get cut off due to buffer end, we\n        // need to be able to just pick right up with the next call where we left off.\n        // The caller decides when to reset things by sending the appropriate\n        // instruction(s).\n\n\t\tlet state = window.vuguRenderState || {};\n\t\twindow.vuguRenderState = state;\n\n\t\tconsole.log(\"vuguRender called\", buffer);\n\n\t\tlet bufferView = new DataView(buffer.buffer, buffer.byteOffset, buffer.byteLength);\n\n        var decoder = new Decoder(bufferView, 0);\n        \n        // state.refMap = state.refMap || {};\n        // state.curRef = state.curRef || \"\"; // current reference number (as a hex string)\n        // state.curRefEl = state.curRefEl || null; // current reference element\n        // state.elStack = state.elStack || []; // stack of elements as we traverse the DOM tree\n\n        // mount point element\n        state.mountPointEl = state.mountPointEl || null; \n\n        // currently selected element\n        state.el = state.el || null;\n\n        // specifies a \"next\" move for the current element, if used it must be followed by\n        // one of opcodeSetElement, opcodeSetText, opcodeSetComment, which will create/replace/use existing\n        // the element and put it in \"el\".  The point is this allow us to select nodes that may\n        // not exist yet, knowing that the next call will specify what that node is.  It's more complex here\n        // but makes it easier to generate instructions while walking a DOM tree.\n        // Value is one of \"first_child\", \"next_sibling\"\n        // (Parents always exist and so doesn't use this mechanism.)\n        state.nextElMove = state.nextElMove || null;\n\n        // keeps track of attributes that are being set on an element, we can remove any extras\n        state.elAttrNames = state.elAttrNames || {};\n\n        instructionLoop: while (true) {\n\n\t\t\tlet opcode = decoder.readUint8();\n\n            switch (opcode) {\n\n                case opcodeEnd: {\n                    break instructionLoop;\n                }\n    \n                // case opcodeClearRefmap:\n                //     state.refMap = {};\n                //     state.curRef = \"\";\n                //     state.curRefEl = null;\n                //     break;\n\n                case opcodeClearEl: {\n                    state.el = null;\n                    state.nextElMove = null;\n                    break;\n                }\n        \n                // case opcodeSetHTMLRef:\n                //     var refstr = decoder.readRefToString();\n                //     state.refMap[refstr] = document.querySelector(\"html\");\n                //     break;\n\n                // case opcodeSelectRef:\n                //     var refstr = decoder.readRefToString();\n                //     state.curRef = refstr;\n                //     state.curRefEl = state.refMap[refstr];\n                //     if (!state.curRefEl) {\n                //         throw \"opcodeSelectRef: refstr does not exist - \" + refstr;\n                //     }\n                //     break;\n\n                case opcodeSetAttrStr: {\n                    let el = state.el;\n                    if (!el) {\n                        return \"opcodeSetAttrStr: no current reference\";\n                    }\n                    let attrName = decoder.readString();\n                    let attrValue = decoder.readString();\n                    el.setAttribute(attrName, attrValue);\n                    state.elAttrNames[attrName] = true;\n                    // console.log(\"setting attr\", attrName, attrValue, el)\n                    break;\n                }\n\n                case opcodeSelectMountPoint: {\n                    \n                    state.elAttrNames = {}; // reset attribute list\n\n                    // select mount point using selector or if it was done earlier re-use the one from before\n                    let selector = decoder.readString();\n                    let nodeName = decoder.readString();\n                    // console.log(\"GOT HERE selector,nodeName = \", selector, nodeName);\n                    // console.log(\"state.mountPointEl\", state.mountPointEl);\n                    if (state.mountPointEl) {\n                        state.el = state.mountPointEl;\n                        // state.elStack.push(state.mountPointEl);\n                    } else {\n                        let el = document.querySelector(selector);\n                        if (!el) {\n                            throw \"mount point selector not found: \" + selector;\n                        }\n                        state.mountPointEl = el;\n                        // state.elStack.push(el);\n                        state.el = el;\n                    }\n\n                    let el = state.el;\n\n                    // make sure it's the right element name and replace if not\n                    if (el.nodeName.toUpperCase() != nodeName.toUpperCase()) {\n\n                        var newEl = document.createElement(nodeName);\n                        el.parentNode.replaceChild(newEl, el);\n\n                        state.mountPointEl = newEl;\n                        el = newEl;\n\n                    }\n\n                    state.el = el;\n\n                    state.nextElMove = null;\n\n                    break;\n                }\n\n                // case opcodePicardFirstChild: {\n\n            \t// \tlet nodeType = decoder.readUint8();\n                //     let data = decoder.readString();\n\n                //     let oldFirstChildEl = state.el.firstChild;\n\n                //     let newFirstChildEl = null;\n\n                //     let needsCreate = true;\n                //     if (oldFirstChildEl) {\n                //         // node types from Go are https://godoc.org/golang.org/x/net/html#NodeType\n                //         // whereas node types in DOM are https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeType\n\n                //         // text\n                //         if (nodeType == 1 && oldFirstChildEl.nodeType == 3) {\n                //             needsCreate = false;\n                //         } else \n                //         // element\n                //         if (nodeType == 3 && oldFirstChildEl.nodeType == 1) {\n                //             needsCreate = false;\n                //         } else \n                //         // comment\n                //         if (nodeType == 4 && oldFirstChildEl.nodeType == 8) {\n                //             needsCreate = false;\n                //         }\n\n                //     }\n\n                //     if (needsCreate) {\n\n                //         switch (nodeType) {\n                //             case 1: {\n                //                 newFirstChildEl = document.createTextNode(data);\n                //                 break;\n                //             }\n                //             case 3: {\n                //                 newFirstChildEl = document.createElement(data);\n                //                 break;\n                //             }\n                //             case 4: {\n                //                 newFirstChildEl = document.createComment(data);\n                //                 break;\n                //             }\n                //         }\n    \n                //     }\n\n                //     if (newFirstChildEl) {\n                //         if (oldFirstChildEl) {\n                //             state.el.replaceChild(newFirstChildEl, oldFirstChildEl);\n                //         } else {\n                //             state.el.appendChild(newFirstChildEl);\n                //         }\n                //         state.el = newFirstChildEl;\n                //     } else {\n                //         state.el = oldFirstChildEl;\n                //     }\n\n                //     break;\n                // }\n\n                // case opcodePicardFirstChildElement: {\n                //     // ensure an element first child and select\n\n                //     let el = state.el;\n                //     let nextEl = el.firstChild;\n                //     if (!nextEl) {\n                //         nextEl = \n                //     }\n                //     state.el = el;\n\n                //     break;\n                // }\n\n                // case opcodePicardFirstChildText: {\n                //     // ensure a text first child and select\n                //     break;\n                // }\n\n                // case opcodePicardFirstChildComment: {\n                //     // ensure a comment first child and select\n                //     break;\n                // }\n\n                // remove any elements for the current element that we didn't just set\n                case opcodeRemoveOtherAttrs: {\n\n                    if (!state.el) {\n                        throw \"no element selected\";\n                    }\n\n                    if (state.nextElMove) {\n                        throw \"cannot call opcodeRemoveOtherAttrs when nextElMove is set\";\n                    }\n\n                    // build a list of attribute names to remove\n                    let rmAttrNames = [];\n                    for (let i = 0; i < state.el.attributes.length; i++) {\n                        if (!state.elAttrNames[state.el.attributes[i].name]) {\n                            rmAttrNames.push(state.el.attributes[i].name);\n                        }\n                    }\n\n                    // remove them\n                    for (let i = 0; i < rmAttrNames.length; i++) {\n                        state.el.attributes.removeNamedItem(rmAttrNames[i]);\n                    }\n\n                    break;\n                }\n\n                // move node selection to parent\n                case opcodeMoveToParent: {\n\n                    // if first_child is next move then we just unset this\n                    if (state.nextElMove == \"first_child\") {\n                        state.nextElMove = null;\n                    } else {\n                        // otherwise we actually move and also reset nextElMove\n                        state.el = state.el.parentNode;\n                        state.nextElMove = null;\n                    }\n\n                    break;\n                }\n\n                // move node selection to first child (doesn't have to exist)\n                case opcodeMoveToFirstChild: {\n\n                    // if a next move already set, then we need to execute it before we can do this\n                    if (state.nextElMove) {\n                        if (state.nextElMove == \"first_child\") {\n                            state.el = state.el.firstChild;\n                            if (!state.el) { throw \"unable to find state.el.firstChild\"; }\n                        } else if (state.nextElMove == \"next_sibling\") {\n                            state.el = state.el.nextSibling;\n                            if (!state.el) { throw \"unable to find state.el.nextSibling\"; }\n                        }\n                        state.nextElMove = null;\n                    }\n\n                    if (!state.el) { throw \"must have current selection to use opcodeMoveToFirstChild\"; }\n                    state.nextElMove = \"first_child\";\n\n                    break;\n                }\n                \n                // move node selection to next sibling (doesn't have to exist)\n                case opcodeMoveToNextSibling: {\n\n                    // if a next move already set, then we need to execute it before we can do this\n                    if (state.nextElMove) {\n                        if (state.nextElMove == \"first_child\") {\n                            state.el = state.el.firstChild;\n                            if (!state.el) { throw \"unable to find state.el.firstChild\"; }\n                        } else if (state.nextElMove == \"next_sibling\") {\n                            state.el = state.el.nextSibling;\n                            if (!state.el) { throw \"unable to find state.el.nextSibling\"; }\n                        }\n                        state.nextElMove = null;\n                    }\n\n                    if (!state.el) { throw \"must have current selection to use opcodeMoveToNextSibling\"; }\n                    state.nextElMove = \"next_sibling\";\n\n                    break;\n                }\n                \n                // assign current selected node as an element of the specified type\n                case opcodeSetElement: {\n\n                    let nodeName = decoder.readString();\n\n                    state.elAttrNames = {};\n\n                    // handle nextElMove cases\n\n                    if (state.nextElMove == \"first_child\") {\n                        state.nextElMove = null;\n                        let newEl = state.el.firstChild;\n                        if (!newEl) {\n                            newEl = document.createElement(nodeName);\n                            state.el.appendChild(newEl);\n                            state.el = newEl;\n                            break; // we're done here, since we just created the right element\n                        }\n                    } else if (state.nextElMove == \"next_sibling\") {\n                        state.nextElMove = null;\n                        let newEl = state.el.nextSibling;\n                        if (!newEl) {\n                            newEl = document.createElement(nodeName);\n                            // console.log(\"HERE1\", state.el);\n                            // state.el.insertAdjacentElement(newEl, 'afterend');\n                            state.el.parentNode.appendChild(newEl);\n                            state.el = newEl;\n                            break; // we're done here, since we just created the right element\n                        }\n                    } else if (state.nextElMove) {\n                        throw \"bad state.nextElMove value: \" + state.nextElMove;\n                    }\n\n                    // if we get here we need to verify that state.el is in fact an element of the right type\n                    // and replace if not\n\n                    if (state.el.nodeType != 1 || state.el.nodeName != nodeName) {\n\n                        let newEl = document.createElement(nodeName);\n                        state.el.parentNode.replaceChild(newEl, state.el);\n                        state.el = newEl;\n\n                    }\n\n                    break;\n                }\n\n                // assign current selected node as text with specified content\n                case opcodeSetText: {\n\n                    let content = decoder.readString();\n\n                    // handle nextElMove cases\n\n                    if (state.nextElMove == \"first_child\") {\n                        state.nextElMove = null;\n                        let newEl = state.el.firstChild;\n                        if (!newEl) {\n                            let newEl = document.createTextNode(content);\n                            state.el.appendChild(newEl);\n                            state.el = newEl;\n                            break; // we're done here, since we just created the right element\n                        }\n                    } else if (state.nextElMove == \"next_sibling\") {\n                        state.nextElMove = null;\n                        let newEl = state.el.nextSibling;\n                        if (!newEl) {\n                            let newEl = document.createTextNode(content);\n                            // state.el.insertAdjacentElement(newEl, 'afterend');\n                            state.el.parentNode.appendChild(newEl);\n                            state.el = newEl;\n                            break; // we're done here, since we just created the right element\n                        }\n                    } else if (state.nextElMove) {\n                        throw \"bad state.nextElMove value: \" + state.nextElMove;\n                    }\n\n                    // if we get here we need to verify that state.el is in fact a node of the right type\n                    // and with right content and replace if not\n\n                    if (state.el.nodeType != 3) {\n\n                        let newEl = document.createTextNode(content);\n                        state.el.parentNode.replaceChild(newEl, state.el);\n                        state.el = newEl;\n\n                    } else {\n                        state.el.textContent = content;\n                    }\n\n                    break;\n                }\n\n                // assign current selected node as comment with specified content\n                case opcodeSetComment: {\n                    \n                    let content = decoder.readString();\n\n                    // handle nextElMove cases\n\n                    if (state.nextElMove == \"first_child\") {\n                        state.nextElMove = null;\n                        let newEl = state.el.firstChild;\n                        if (!newEl) {\n                            let newEl = document.createComment(content);\n                            state.el.appendChild(newEl);\n                            state.el = newEl;\n                            break; // we're done here, since we just created the right element\n                        }\n                    } else if (state.nextElMove == \"next_sibling\") {\n                        state.nextElMove = null;\n                        let newEl = state.el.nextSibling;\n                        if (!newEl) {\n                            let newEl = document.createComment(content);\n                            // state.el.insertAdjacentElement(newEl, 'afterend');\n                            state.el.parentNode.appendChild(newEl);\n                            state.el = newEl;\n                            break; // we're done here, since we just created the right element\n                        }\n                    } else if (state.nextElMove) {\n                        throw \"bad state.nextElMove value: \" + state.nextElMove;\n                    }\n\n                    // if we get here we need to verify that state.el is in fact a node of the right type\n                    // and with right content and replace if not\n\n                    if (state.el.nodeType != 8) {\n\n                        let newEl = document.createComment(content);\n                        state.el.parentNode.replaceChild(newEl, state.el);\n                        state.el = newEl;\n\n                    } else {\n                        state.el.textContent = content;\n                    }\n\n                    break;\n                }\n\n                // remove all event listeners from currently selected element\n                case opcodeClearEventListeners: {\n                    break;\n                }\n\n                // assign event listener to currently selected element\n                case opcodeSetEventListener: {\n                    break;\n                }\n            \n                // case opcodeSelectParent: {\n                //     // select parent\n                //     state.el = state.el.parentNode;\n                //     break;\n                // }\n\n                default: {\n                    console.error(\"found invalid opcode\", opcode);\n                    return;\n                }\n            }\n\n\t\t}\n\n\t}\n\n})()\n"
