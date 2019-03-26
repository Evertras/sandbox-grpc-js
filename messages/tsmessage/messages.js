/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

$root.data = (function() {

    /**
     * Namespace data.
     * @exports data
     * @namespace
     */
    var data = {};

    data.EchoMessage = (function() {

        /**
         * Properties of an EchoMessage.
         * @memberof data
         * @interface IEchoMessage
         * @property {string|null} [greeting] EchoMessage greeting
         */

        /**
         * Constructs a new EchoMessage.
         * @memberof data
         * @classdesc Represents an EchoMessage.
         * @implements IEchoMessage
         * @constructor
         * @param {data.IEchoMessage=} [properties] Properties to set
         */
        function EchoMessage(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * EchoMessage greeting.
         * @member {string} greeting
         * @memberof data.EchoMessage
         * @instance
         */
        EchoMessage.prototype.greeting = "";

        /**
         * Creates a new EchoMessage instance using the specified properties.
         * @function create
         * @memberof data.EchoMessage
         * @static
         * @param {data.IEchoMessage=} [properties] Properties to set
         * @returns {data.EchoMessage} EchoMessage instance
         */
        EchoMessage.create = function create(properties) {
            return new EchoMessage(properties);
        };

        /**
         * Encodes the specified EchoMessage message. Does not implicitly {@link data.EchoMessage.verify|verify} messages.
         * @function encode
         * @memberof data.EchoMessage
         * @static
         * @param {data.IEchoMessage} message EchoMessage message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        EchoMessage.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.greeting != null && message.hasOwnProperty("greeting"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.greeting);
            return writer;
        };

        /**
         * Encodes the specified EchoMessage message, length delimited. Does not implicitly {@link data.EchoMessage.verify|verify} messages.
         * @function encodeDelimited
         * @memberof data.EchoMessage
         * @static
         * @param {data.IEchoMessage} message EchoMessage message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        EchoMessage.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an EchoMessage message from the specified reader or buffer.
         * @function decode
         * @memberof data.EchoMessage
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {data.EchoMessage} EchoMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        EchoMessage.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.data.EchoMessage();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.greeting = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an EchoMessage message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof data.EchoMessage
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {data.EchoMessage} EchoMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        EchoMessage.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an EchoMessage message.
         * @function verify
         * @memberof data.EchoMessage
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        EchoMessage.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.greeting != null && message.hasOwnProperty("greeting"))
                if (!$util.isString(message.greeting))
                    return "greeting: string expected";
            return null;
        };

        /**
         * Creates an EchoMessage message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof data.EchoMessage
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {data.EchoMessage} EchoMessage
         */
        EchoMessage.fromObject = function fromObject(object) {
            if (object instanceof $root.data.EchoMessage)
                return object;
            var message = new $root.data.EchoMessage();
            if (object.greeting != null)
                message.greeting = String(object.greeting);
            return message;
        };

        /**
         * Creates a plain object from an EchoMessage message. Also converts values to other types if specified.
         * @function toObject
         * @memberof data.EchoMessage
         * @static
         * @param {data.EchoMessage} message EchoMessage
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        EchoMessage.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.greeting = "";
            if (message.greeting != null && message.hasOwnProperty("greeting"))
                object.greeting = message.greeting;
            return object;
        };

        /**
         * Converts this EchoMessage to JSON.
         * @function toJSON
         * @memberof data.EchoMessage
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        EchoMessage.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return EchoMessage;
    })();

    data.Echo = (function() {

        /**
         * Constructs a new Echo service.
         * @memberof data
         * @classdesc Represents an Echo
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function Echo(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (Echo.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = Echo;

        /**
         * Creates new Echo service using the specified rpc implementation.
         * @function create
         * @memberof data.Echo
         * @static
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         * @returns {Echo} RPC service. Useful where requests and/or responses are streamed.
         */
        Echo.create = function create(rpcImpl, requestDelimited, responseDelimited) {
            return new this(rpcImpl, requestDelimited, responseDelimited);
        };

        /**
         * Callback as used by {@link data.Echo#sayHello}.
         * @memberof data.Echo
         * @typedef SayHelloCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {data.EchoMessage} [response] EchoMessage
         */

        /**
         * Calls SayHello.
         * @function sayHello
         * @memberof data.Echo
         * @instance
         * @param {data.IEchoMessage} request EchoMessage message or plain object
         * @param {data.Echo.SayHelloCallback} callback Node-style callback called with the error, if any, and EchoMessage
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(Echo.prototype.sayHello = function sayHello(request, callback) {
            return this.rpcCall(sayHello, $root.data.EchoMessage, $root.data.EchoMessage, request, callback);
        }, "name", { value: "SayHello" });

        /**
         * Calls SayHello.
         * @function sayHello
         * @memberof data.Echo
         * @instance
         * @param {data.IEchoMessage} request EchoMessage message or plain object
         * @returns {Promise<data.EchoMessage>} Promise
         * @variation 2
         */

        return Echo;
    })();

    return data;
})();

module.exports = $root;
