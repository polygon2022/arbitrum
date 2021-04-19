/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Signer } from 'ethers'
import { Provider, TransactionRequest } from '@ethersproject/providers'
import { Contract, ContractFactory, Overrides } from '@ethersproject/contracts'

import type { Validator } from '../Validator'

export class Validator__factory extends ContractFactory {
  constructor(signer?: Signer) {
    super(_abi, _bytecode, signer)
  }

  deploy(overrides?: Overrides): Promise<Validator> {
    return super.deploy(overrides || {}) as Promise<Validator>
  }
  getDeployTransaction(overrides?: Overrides): TransactionRequest {
    return super.getDeployTransaction(overrides || {})
  }
  attach(address: string): Validator {
    return super.attach(address) as Validator
  }
  connect(signer: Signer): Validator__factory {
    return super.connect(signer) as Validator__factory
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): Validator {
    return new Contract(address, _abi, signerOrProvider) as Validator
  }
}

const _abi = [
  {
    inputs: [],
    stateMutability: 'nonpayable',
    type: 'constructor',
  },
  {
    inputs: [
      {
        internalType: 'bytes',
        name: 'data',
        type: 'bytes',
      },
      {
        internalType: 'address',
        name: 'destination',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'amount',
        type: 'uint256',
      },
    ],
    name: 'executeTransaction',
    outputs: [],
    stateMutability: 'payable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'bytes[]',
        name: 'data',
        type: 'bytes[]',
      },
      {
        internalType: 'address[]',
        name: 'destination',
        type: 'address[]',
      },
      {
        internalType: 'uint256[]',
        name: 'amount',
        type: 'uint256[]',
      },
    ],
    name: 'executeTransactions',
    outputs: [],
    stateMutability: 'payable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'contract IRollup',
        name: 'rollup',
        type: 'address',
      },
      {
        internalType: 'address payable[]',
        name: 'stakers',
        type: 'address[]',
      },
    ],
    name: 'returnOldDeposits',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'contract IChallenge[]',
        name: 'challenges',
        type: 'address[]',
      },
    ],
    name: 'timeoutChallenges',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
]

const _bytecode =
  '0x608060405234801561001057600080fd5b50600080546001600160a01b03191633179055610734806100326000396000f3fe60806040526004361061003f5760003560e01c806372f458661461004457806381aac2d914610059578063944f449514610079578063ce1d571f14610099575b600080fd5b61005761005236600461048c565b6100ac565b005b34801561006557600080fd5b50610057610074366004610522565b6101c1565b34801561008557600080fd5b506100576100943660046105e7565b61029b565b6100576100a7366004610562565b610381565b6000546001600160a01b031633146100df5760405162461bcd60e51b81526004016100d69061067b565b60405180910390fd5b8460005b818110156101b75760008686838181106100f957fe5b905060200201602081019061010e9190610469565b6001600160a01b031685858481811061012357fe5b905060200201358a8a8581811061013657fe5b9050602002810190610148919061069f565b60405161015692919061063a565b60006040518083038185875af1925050503d8060008114610193576040519150601f19603f3d011682016040523d82523d6000602084013e610198565b606091505b50509050806101ae576040513d806000833e8082fd5b506001016100e3565b5050505050505050565b8060005b81811015610295578383828181106101d957fe5b90506020020160208101906101ee9190610469565b6001600160a01b03166370dea79a6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561022857600080fd5b505af1925050508015610239575060015b61028d573d808015610267576040519150601f19603f3d011682016040523d82523d6000602084013e61026c565b606091505b50805161028b5760405162461bcd60e51b81526004016100d69061065e565b505b6001016101c5565b50505050565b8060005b8181101561037a57846001600160a01b0316637427be518585848181106102c257fe5b90506020020160208101906102d79190610469565b6040518263ffffffff1660e01b81526004016102f3919061064a565b600060405180830381600087803b15801561030d57600080fd5b505af192505050801561031e575060015b610372573d80801561034c576040519150601f19603f3d011682016040523d82523d6000602084013e610351565b606091505b5080516103705760405162461bcd60e51b81526004016100d69061065e565b505b60010161029f565b5050505050565b6000546001600160a01b031633146103ab5760405162461bcd60e51b81526004016100d69061067b565b6000826001600160a01b03168286866040516103c892919061063a565b60006040518083038185875af1925050503d8060008114610405576040519150601f19603f3d011682016040523d82523d6000602084013e61040a565b606091505b505090508061037a576040513d806000833e8082fd5b60008083601f840112610431578182fd5b50813567ffffffffffffffff811115610448578182fd5b602083019150836020808302850101111561046257600080fd5b9250929050565b60006020828403121561047a578081fd5b8135610485816106e6565b9392505050565b600080600080600080606087890312156104a4578182fd5b863567ffffffffffffffff808211156104bb578384fd5b6104c78a838b01610420565b909850965060208901359150808211156104df578384fd5b6104eb8a838b01610420565b90965094506040890135915080821115610503578384fd5b5061051089828a01610420565b979a9699509497509295939492505050565b60008060208385031215610534578182fd5b823567ffffffffffffffff81111561054a578283fd5b61055685828601610420565b90969095509350505050565b60008060008060608587031215610577578384fd5b843567ffffffffffffffff8082111561058e578586fd5b81870188601f82011261059f578687fd5b80359250818311156105af578687fd5b8860208483010111156105c0578687fd5b60209081019650919450508501356105d7816106e6565b9396929550929360400135925050565b6000806000604084860312156105fb578283fd5b8335610606816106e6565b9250602084013567ffffffffffffffff811115610621578283fd5b61062d86828701610420565b9497909650939450505050565b6000828483379101908152919050565b6001600160a01b0391909116815260200190565b60208082526003908201526247415360e81b604082015260600190565b6020808252600a908201526927a7262cafa7aba722a960b11b604082015260600190565b6000808335601e198436030181126106b5578283fd5b8084018035925067ffffffffffffffff8311156106d0578384fd5b6020019250503681900382131561046257600080fd5b6001600160a01b03811681146106fb57600080fd5b5056fea26469706673582212206d7614448d9173570ce2119f80160448fbe9c28b4a2bed5f21ed0702723fa24e64736f6c634300060b0033'