# frozen_string_literal: true

RSpec.describe CashesController, type: :controller do
  describe '#collection' do
    context do
      before { subject.instance_variable_set :@collection, :collection }

      its(:collection) { should eq :collection }
    end

    context do
      before { allow(subject).to receive(:params).and_return(:params) }

      before { allow(Cashes::GetCollectionService).to receive(:call).with(:params).and_return(:collection) }

      its(:collection) { should eq :collection }
    end
  end

  describe '#resource' do
    context do
      before { subject.instance_variable_set :@resource, :resource }

      its(:resource) { should eq :resource }
    end

    context do
      before { allow(subject).to receive(:params).and_return(:params) }

      before { allow(Cashes::GetResourceService).to receive(:call).with(:params).and_return(:resource) }

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
