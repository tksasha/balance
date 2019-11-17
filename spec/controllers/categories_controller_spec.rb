# frozen_string_literal: true

RSpec.describe CategoriesController, type: :controller do
  describe '#relation' do
    before { expect(Category).to receive(:order).with(:income).and_return(:relation) }

    its(:relation) { should eq :relation }
  end

  describe '#collection' do
    context do
      before { subject.instance_variable_set :@collection, :collection }

      its(:collection) { should eq :collection }
    end

    context do
      let(:relation) { double }

      let(:params) { { currency: 'usd' } }

      before { expect(subject).to receive(:relation).and_return(relation) }

      before { allow(subject).to receive(:params).and_return(params) }

      before { expect(CategorySearcher).to receive(:search).with(relation, params).and_return(:collection) }

      its(:collection) { should eq :collection }
    end
  end

  describe '#resource' do
    context do
      before { subject.instance_variable_set :@resource, :resource }

      its(:resource) { should eq :resource }
    end

    context do
      before { expect(subject).to receive(:params).and_return(id: 27) }

      before { expect(Category).to receive(:find).with(27).and_return(:resource) }

      its(:resource) { should eq :resource }
    end
  end

  describe '#initialize_resource' do
    let(:params) { { currency: 'usd', foo: 'bar' } }

    before { expect(subject).to receive(:params).and_return(params) }

    before { expect(Category).to receive(:new).with(currency: 'usd').and_return(:resource) }

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

    before { expect(subject).to receive(:params).and_return(params) }

    its(:resource_params) { should eq params.require(:category).permit! }
  end

  describe '#build_resource' do
    before { expect(subject).to receive(:resource_params).and_return(:resource_params) }

    before { expect(Category).to receive(:new).with(:resource_params).and_return(:resource) }

    before { subject.send :build_resource }

    its(:resource) { should eq :resource }
  end
end
