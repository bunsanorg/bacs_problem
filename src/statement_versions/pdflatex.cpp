#include "pdflatex.hpp"

#include <bunsan/config/cast.hpp>
#include <bunsan/filesystem/fstream.hpp>
#include <bunsan/filesystem/operations.hpp>
#include <bunsan/pm/index.hpp>

#include <boost/property_tree/ini_parser.hpp>

namespace bacs{namespace problem{namespace statement_versions
{
    const bool pdflatex::factory_reg_hook = statement::version::register_new("pdflatex",
        [](const boost::filesystem::path &location,
           const boost::property_tree::ptree &config)
        {
            statement::version_ptr tmp(new pdflatex(location, config));
            return tmp;
        });

    namespace
    {
        boost::filesystem::path get_root(const boost::filesystem::path &source)
        {
            return source.parent_path();
        }
        boost::filesystem::path get_source(const boost::filesystem::path &source)
        {
            return source.filename();
        }
        boost::filesystem::path get_target(const boost::filesystem::path &source)
        {
            return source.stem().string() + ".pdf";
        }
    }

    pdflatex::pdflatex(const boost::filesystem::path &/*location*/,
                       const boost::property_tree::ptree &config):
        version(config.get<std::string>("info.lang"), "pdf"),
        m_root(get_root(config.get<std::string>("build.source"))),
        m_source(get_source(config.get<std::string>("build.source"))),
        m_target(get_target(config.get<std::string>("build.source"))) {}

    void pdflatex::make_package(const boost::filesystem::path &destination,
                                const bunsan::pm::entry &/*package*/,
                                const bunsan::pm::entry &resources_package) const
    {
        try
        {
            bunsan::filesystem::reset_dir(destination);
            bunsan::pm::index index;
            index.source.import.source.insert(std::make_pair(".", "bacs/system/statement/pdflatex"));
            index.source.import.source.insert(std::make_pair("src", resources_package));
            index.source.self.insert(std::make_pair(".", "src"));
            index.package.self.insert(std::make_pair(".", "pkg"));
            boost::filesystem::create_directory(destination / "src");
            bunsan::filesystem::ofstream fout(destination / "src" / "source.cmake");
            BUNSAN_FILESYSTEM_FSTREAM_WRAP_BEGIN(fout)
            {
                fout << "set(root " << m_root << ")\n";
                fout << "set(source " << m_source << ")\n";
                fout << "set(target " << m_target << ")\n";
            }
            BUNSAN_FILESYSTEM_FSTREAM_WRAP_END(fout)
            fout.close();
            boost::filesystem::create_directory(destination / "pkg");
            manifest statement_manifest;
            statement_manifest.version.lang = lang();
            statement_manifest.version.format = format();
            statement_manifest.data.index = m_target;
            boost::property_tree::write_ini((destination / "pkg" / manifest_path).string(),
                                            bunsan::config::save<boost::property_tree::ptree>(statement_manifest));
            index.save(destination / "index");
        }
        catch (std::exception &)
        {
            BOOST_THROW_EXCEPTION(statement_version_make_package_error() <<
                                  statement_version_make_package_error::destination(destination) <<
                                  //statement_version_make_package_error::package(package) <<
                                  statement_version_make_package_error::resources_package(resources_package) <<
                                  bunsan::enable_nested_current());
        }
    }
}}}
