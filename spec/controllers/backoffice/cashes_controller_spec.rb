# frozen_string_literal: true

RSpec.describe Backoffice::CashesController, type: :controller do
  describe '#cashes' do
    let(:cashes) { double }

    before { allow(Cash).to receive(:order).with(:name).and_return(cashes) }

    its(:cashes) { should eq cashes }
  end

  describe '#collection' do
    let(:collection) { double }

    context do
      before { subject.instance_variable_set :@collection, collection }

      its(:collection) { should eq collection }
    end

    context do
      let(:cashes) { double }

      let(:params) { double }

      before { allow(subject).to receive(:cashes).and_return(cashes) }

      before { allow(subject).to receive(:params).and_return(params) }

      before { allow(CashSearcher).to receive(:search).with(cashes, params).and_return(collection) }

      its(:collection) { should eq collection }
    end
  end

  describe '#resource' do
    context do
      before { subject.instance_variable_set :@resource, :resource }

      its(:resource) { should eq :resource }
    end

    context do
      before { allow(subject).to receive(:params).and_return(id: 11) }

      before { allow(Cash).to receive(:find).with(11).and_return(:resource) }

      its(:resource) { should eq :resource }
    end
  end

  describe '#initialize_resource' do
    let(:params) { { currency: 'usd', foo: 'foo' } }

    before { allow(subject).to receive(:params).and_return(params) }

    before { allow(Cash).to receive(:new).with(currency: 'usd').and_return(:resource) }

    before { subject.send :initialize_resource }

    its(:resource) { should eq :resource }
  end

  describe '#resource_params' do
    let :params do
      acp \
        cash: {
          name: nil,
          formula: nil,
          currency: nil,
        }
    end

    before { allow(subject).to receive(:params).and_return(params) }

    its(:resource_params) { should eq params.require(:cash).permit! }
  end

  describe '#build_resource' do
    before { allow(subject).to receive(:resource_params).and_return(:resource_params) }

    before { allow(Cash).to receive(:new).with(:resource_params).and_return(:resource) }

    before { subject.send :build_resource }

    its(:resource) { should eq :resource }
  end
end
