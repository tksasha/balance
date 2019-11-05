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

      before { expect(subject).to receive(:params).twice.and_return(params) }

      before { expect(Consolidation).to receive(:includes).with(:category).and_return(relation) }

      before do
        #
        # ConsolidationSearcher.search(relation, params).decorate(context: params) -> :collection
        #
        expect(ConsolidationSearcher).to receive(:search).with(relation, params) do
          double.tap { |a| expect(a).to receive(:decorate).with(context: params).and_return(:collection) }
        end
      end

      its(:collection) { should eq :collection }
    end
  end
end
