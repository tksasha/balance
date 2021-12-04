# frozen_string_literal: true

RSpec.describe ItemsController, type: :controller do
  describe '#collection' do
    context do
      before { subject.instance_variable_set :@collection, :collection }

      its(:collection) { should eq :collection }
    end

    context do
      before { allow(subject).to receive(:params).and_return(:params) }

      before { allow(Items::GetCollectionService).to receive(:call).with(:params).and_return(:collection) }

      its(:collection) { should eq :collection }
    end
  end

  describe '#resource' do
    context do
      before { subject.instance_variable_set :@resource, :resource }

      its(:resource) { should eq :resource }
    end

    context do
      before { allow(subject).to receive(:params).and_return(id: 26) }

      before { allow(Item).to receive(:find).with(26).and_return(:resource) }

      its(:resource) { should eq :resource }
    end
  end

  describe '#resource_params' do
    let :params do
      acp \
        item: {
          date: nil,
          formula: nil,
          category_id: nil,
          description: nil,
          currency: nil,
        }
    end

    before { allow(subject).to receive(:params).and_return(params) }

    its(:resource_params) { should eq params.require(:item).permit! }
  end

  describe '#build_resource' do
    before { allow(subject).to receive(:resource_params).and_return(:resource_params) }

    before { allow(Item).to receive(:new).with(:resource_params).and_return(:resource) }

    before { subject.send :build_resource }

    its(:resource) { should eq :resource }
  end
end
