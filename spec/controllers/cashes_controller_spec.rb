# frozen_string_literal: true

RSpec.describe CashesController, type: :controller do
  describe '#cashes' do
    before { allow(Cash).to receive(:order).with(:name).and_return(:cashes) }

    its(:cashes) { should eq :cashes }
  end

  describe '#collection' do
    context do
      before { subject.instance_variable_set :@collection, :collection }

      its(:collection) { should eq :collection }
    end

    context do
      let(:cashes) { double }

      let(:params) { double }

      before { allow(subject).to receive(:cashes).and_return(cashes) }

      before { allow(subject).to receive(:params).and_return(params) }

      before { allow(CashSearcher).to receive(:search).with(cashes, params).and_return(:collection) }

      its(:collection) { should eq :collection }
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
end
