require 'rails_helper'

RSpec.describe ConsolidationsController, type: :controller do
  describe '#date' do
    before { expect(subject).to receive(:params).and_return(:params) }

    before { expect(DateFactory).to receive(:build).with(:params).and_return(:date) }

    its(:date) { should eq :date }
  end

  describe '#collection' do
    context do
      before { subject.instance_variable_set :@collection, :collection }

      its(:collection) { should eq :collection }
    end

    context do
      before { expect(subject).to receive(:date).twice.and_return(:date) }

      before { expect(Consolidation).to receive(:includes).with(:category).and_return(:relation) }

      before do
        #
        # ConsolidationSearcher.search(:relation, date: :date).decorate(context: { date: :date }) -> :collection
        #
        expect(ConsolidationSearcher).to receive(:search).with(:relation, date: :date) do
          double.tap { |a| expect(a).to receive(:decorate).with(context: { date: :date }).and_return(:collection) }
        end
      end

      its(:collection) { should eq :collection }
    end
  end

  it_behaves_like :index, format: :js
end
