# frozen_string_literal: true

RSpec.describe Items::DestroyService do
  subject { described_class.new params }

  let(:params) { { id: 25 } }

  describe '#item' do
    context do
      before { subject.instance_variable_set :@item, :item }

      its(:item) { should eq :item }
    end

    context do
      before { allow(Item).to receive(:find).with(25).and_return(:item) }

      its(:item) { should eq :item }
    end
  end

  describe '#call' do
    let(:item) { stub_model Item }

    before { allow(subject).to receive(:item).and_return(item) }

    before { allow(item).to receive(:destroy).and_return(true) }

    its(:call) { should be_success }

    its('call.object') { should eq item }
  end
end
