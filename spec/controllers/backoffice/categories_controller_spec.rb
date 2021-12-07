# frozen_string_literal: true

RSpec.describe Backoffice::CategoriesController, type: :controller do
  describe '#collection' do
    context do
      before { subject.instance_variable_set :@collection, :collection }

      its(:collection) { should eq :collection }
    end

    context do
      before { allow(subject).to receive(:params).and_return(:params) }

      before { allow(Categories::GetCollectionService).to receive(:call).with(:params).and_return(:collection) }

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

      before { allow(Categories::GetResultService).to receive(:call).with(:action_name, :params).and_return(:result) }

      its(:result) { should eq :result }
    end
  end

  describe '#resource_params' do
    let :params do
      acp \
        category: {
          name: nil,
          income: nil,
          visible: nil,
          currency: nil
        }
    end

    before { allow(subject).to receive(:params).and_return(params) }

    its(:resource_params) { should eq params.require(:category).permit! }
  end
end
