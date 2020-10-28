# frozen_string_literal: true

RSpec.describe ItemsController, type: :controller do
  describe '#relation' do
    before do
      #
      # Item
      #   .order(date: :desc)
      #   includes(:category) -> :relation
      #
      allow(Item).to receive(:order).with(date: :desc) do
        double.tap do |a|
          allow(a).to receive(:includes).with(:category).and_return(:relation)
        end
      end
    end

    its(:relation) { should eq :relation }
  end

  describe '#collection' do
    context do
      before { subject.instance_variable_set :@collection, :collection }

      its(:collection) { should eq :collection }
    end

    context do
      let(:params) { double }

      let(:relation) { double }

      before { allow(subject).to receive(:relation).and_return(relation) }

      before { allow(subject).to receive(:params).and_return(params) }

      before { allow(ItemSearcher).to receive(:search).with(relation, params).and_return(:collection) }

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
