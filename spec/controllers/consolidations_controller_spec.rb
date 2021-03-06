# frozen_string_literal: true

RSpec.describe ConsolidationsController, type: :controller do
  describe '#collection' do
    context do
      before { subject.instance_variable_set :@collection, :collection }

      its(:collection) { should eq :collection }
    end

    context do
      let(:relation) { double }

      let(:params) { double }

      before { allow(subject).to receive(:params).and_return(params) }

      before { allow(Consolidation).to receive(:includes).with(:category).and_return(relation) }

      before { allow(ConsolidationSearcher).to receive(:search).with(relation, params).and_return(:collection) }

      its(:collection) { should eq :collection }
    end
  end
end
