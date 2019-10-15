# frozen_string_literal: true

RSpec.describe ItemsController, type: :controller do
  describe '#collection' do
    context do
      before { subject.instance_variable_set :@collection, :collection }

      its(:collection) { should eq :collection }
    end

    context do
      let(:params) { double }

      let(:relation) { double }

      before { expect(subject).to receive(:params).and_return(params) }

      before { expect(Item).to receive(:order).with(date: :desc).and_return(relation) }

      before { expect(ItemSearcher).to receive(:search).with(relation, params).and_return(:collection) }

      its(:collection) { should eq :collection }
    end
  end

  describe '#resource' do
    context do
      before { subject.instance_variable_set :@resource, :resource }

      its(:resource) { should eq :resource }
    end

    context do
      before { expect(subject).to receive(:params).and_return(id: 26) }

      before { expect(Item).to receive(:find).with(26).and_return(:resource) }

      its(:resource) { should eq :resource }
    end
  end

  describe '#build_resource' do
    before { expect(subject).to receive(:resource_params).and_return(:resource_params) }

    before { expect(Item).to receive(:new).with(:resource_params).and_return(:resource) }

    before { subject.send :build_resource }

    its(:resource) { should eq :resource }
  end
end
