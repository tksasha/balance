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

  describe '#resource' do
    context do
      before { subject.instance_variable_set :@resource, :resource }

      its(:resource) { should eq :resource }
    end

    context do
      before { allow(subject).to receive(:params).and_return(id: 27) }

      before { allow(Category).to receive(:find).with(27).and_return(:resource) }

      its(:resource) { should eq :resource }
    end
  end

  describe '#initialize_resource' do
    let(:params) { { currency: 'usd', foo: 'bar' } }

    before { allow(subject).to receive(:params).and_return(params) }

    before { allow(Category).to receive(:new).with(currency: 'usd').and_return(:resource) }

    before { subject.send :initialize_resource }

    its(:resource) { should eq :resource }
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

  describe '#build_resource' do
    before { allow(subject).to receive(:resource_params).and_return(:resource_params) }

    before { allow(Category).to receive(:new).with(:resource_params).and_return(:resource) }

    before { subject.send :build_resource }

    its(:resource) { should eq :resource }
  end
end
