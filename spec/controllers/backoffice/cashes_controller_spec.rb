# frozen_string_literal: true

RSpec.describe Backoffice::CashesController, type: :controller do
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

  describe '#result' do
    context do
      before { subject.instance_variable_set :@result, :result }

      its(:result) { should eq :result }
    end

    context do
      before { allow(subject).to receive(:action_name).and_return(:action_name) }

      before { allow(subject).to receive(:params).and_return(:params) }

      before { allow(Cashes::GetResultService).to receive(:call).with(:action_name, :params).and_return(:result) }

      its(:result) { should eq :result }
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

  pending '#build_resource' do
    before { allow(subject).to receive(:resource_params).and_return(:resource_params) }

    before { allow(Cash).to receive(:new).with(:resource_params).and_return(:resource) }

    before { subject.send :build_resource }

    its(:resource) { should eq :resource }
  end
end
